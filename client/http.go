package client

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

func Initialize(email string, password string) *HTTPSessionManager {
	H := &HTTPSessionManager{
		Email:         email,
		Password:      password,
		LoginResponse: LoginResponse{},
		Client:        resty.New().AddRetryAfterErrorCondition().EnableTrace().SetDisableWarn(true),
		wg:            sync.WaitGroup{},
		mutex:         sync.RWMutex{},
		KeyIndex:      0,
	}
	return H
}

func (h *HTTPSessionManager) APILopin() error {
	var req *resty.Request
	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"email":"%s","password":"%s"}`, h.Email, h.Password)).
		SetResult(&req).
		Post("https://developer.clashofclans.com/api/login")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
	}
	if err := json.Unmarshal(resp.Body(), &h.LoginResponse); err != nil { //login unmarshal
		return err
	}
	return nil
}

func (h *HTTPSessionManager) GetKeys() error {
	var req *resty.Request
	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&req).
		Post("https://developer.clashofclans.com/api/apikey/list")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
	}
	if err := json.Unmarshal(resp.Body(), &h.KeysList); err != nil { //raw json from sc
		return err
	}
	for _, key := range h.KeysList.Keys { //actual key info
		h.RawKeys = append(h.RawKeys, key.Key)
	}
	return nil
}

func (h *HTTPSessionManager) AddKey() error {
	var req *resty.Request
	var keycreationresponse KeyCreationResponse
	desc := fmt.Sprintf("Created on %s", time.Now().Format(time.RFC3339))
	body := fmt.Sprintf(`{"name":"%s","description":"%s", "cidrRanges": ["%s"], "scopes": ["clash"]}`, "coc.go", desc, h.IP)
	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&req).
		Post("https://developer.clashofclans.com/api/apikey/create")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
	}
	if err := json.Unmarshal(resp.Body(), &keycreationresponse); err != nil {
		return err
	}
	h.KeysList.Keys = append(h.KeysList.Keys, keycreationresponse.Key)
	h.RawKeys = append(h.RawKeys, keycreationresponse.Key.Key)
	fmt.Println("Added key with ID", keycreationresponse.Key.ID)
	return nil
}

func (h *HTTPSessionManager) DeleteKey(key Key) error {
	var req *resty.Request
	var keydeletionresponse KeyDeletionResponse
	jsn := fmt.Sprintf(`{"id": "%s"}`, key.ID)
	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsn).
		SetResult(&req).
		Post("https://developer.clashofclans.com/api/apikey/revoke")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
	}
	if err := json.Unmarshal(resp.Body(), &keydeletionresponse); err != nil {
		return err
	}
	h.KeysList.Keys = RemoveKey(h.KeysList.Keys, key)
	return nil
}

func (h *HTTPSessionManager) AddOrDeleteKeysAsNecessary() error {
	errC := make(chan error, len(h.KeysList.Keys))
	err := h.GetIP()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, key := range h.KeysList.Keys {
		h.wg.Add(1)
		go func(key Key) {
			h.mutex.Lock()
			defer h.mutex.Unlock()
			defer h.wg.Done()
			if !Contains(key.Cidrranges, h.IP) {
				err := h.DeleteKey(key)
				if err != nil {
					fmt.Println("FAILED TO DELETE", err.Error())
					errC <- err
					return
				}
				err = h.AddKey()
				if err != nil {
					fmt.Println(err.Error())
					errC <- err
					return
				}
			}
			errC <- nil
		}(key)
		err := <-errC
		if err != nil {
			return err
		}
	}
	h.wg.Wait()

	if len(h.KeysList.Keys) != 10 {
		fmt.Printf("Creating %d additional keys\n", 10-len(h.KeysList.Keys))
		for {
			if len(h.KeysList.Keys) >= 10 { //max limit
				break
			}
			h.wg.Add(1)
			go func() {
				h.mutex.Lock()
				defer h.mutex.Unlock()
				defer h.wg.Done()
				err = h.AddKey()
				if err != nil {
					errC <- err
					return
				}
			}()
		}
		for i := 0; i < len(h.KeysList.Keys); i++ {
			if err := <-errC; err != nil {
				return err
			}
		}
	}
	h.wg.Wait()
	close(errC)
	return nil
}

func (h *HTTPSessionManager) GetIP() error {
	var req *resty.Request
	resp, err := h.Client.R().
		SetResult(&req).
		Get("https://api.ipify.org/")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
	}
	h.IP = string(resp.Body())
	return nil
}

func (h *HTTPSessionManager) ViewData() {
	for _, key := range h.KeysList.Keys {
		fmt.Println("[Key] ", key.Key, "[ID]", key.ID, key.Name, key.Cidrranges, key.Developerid, key.Description)
	}
}

func RemoveKey(keylist []Key, deleteKey Key) []Key {
	var ret []Key
	for _, key := range keylist {
		if deleteKey.ID != key.ID {
			ret = append(ret, key)
		}
	}
	return ret
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
