package client

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
)

func Initialize(credentials ...map[string]string) *HTTPSessionManager {
	var creds []LoginCredential
	for _, credential := range credentials {
		for email, password := range credential {
			creds = append(creds, LoginCredential{Email: email, Password: password})
		}
	}
	H := &HTTPSessionManager{
		Credentials:   creds,
		LoginResponse: LoginResponse{},
		Client:        resty.New().AddRetryAfterErrorCondition().EnableTrace().SetDisableWarn(true),
		WG:            sync.WaitGroup{},
		KeyIndex:      0,
		IsValidKeys:   true,
		cache:         cache.New(time.Second*60, time.Second*60),
	}
	for _, credential := range H.Credentials {
		err := H.APILopin(credential)
		if err != nil {
			fmt.Println(err.Error())
		}
		err = H.GetKeys()
		if err != nil {
			fmt.Println(err.Error())
		}
		err = H.AddOrDeleteKeysAsNecessary(*H.LoginResponse.Developer.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return H
}

func (h *HTTPSessionManager) APILopin(credential LoginCredential) error {
	var req *resty.Request
	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"email":"%s","password":"%s"}`, credential.Email, credential.Password)).
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
	h.RawKeysList = append(h.RawKeysList, h.KeysList.Keys...)
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
	//h.KeysList.Keys = append(h.KeysList.Keys, keycreationresponse.Key)
	h.RawKeysList = append(h.RawKeysList, keycreationresponse.Key)
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
	//h.KeysList.Keys = RemoveKey(h.KeysList.Keys, key)
	h.RawKeysList = RemoveKey(h.RawKeysList, key)
	return nil
}

func (h *HTTPSessionManager) AddOrDeleteKeysAsNecessary(developerID string) error {
	h.IsValidKeys = false
	errC := make(chan error, 10)
	err := h.GetIP()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, key := range h.KeysList.Keys {
		h.WG.Add(1)
		if key.Developerid != developerID {
			continue
		}
		go func(key Key) {
			h.Lock()
			defer h.Unlock()
			defer h.WG.Done()
			if !contains(key.Cidrranges, h.IP) {
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
	h.WG.Wait()
	thisdevskeycount := 0
	for _, key := range h.KeysList.Keys {
		if key.Developerid == developerID {
			thisdevskeycount++
		}
	}
	if thisdevskeycount < 10 {
		fmt.Printf("Creating %d additional keys\n", 10-thisdevskeycount)
		for {
			if thisdevskeycount >= 10 { //max limit
				break
			}
			h.WG.Add(1)
			go func() {
				h.Lock()
				defer h.Unlock()
				defer h.WG.Done()
				err = h.AddKey()
				if err != nil {
					errC <- err
					return
				}
				thisdevskeycount++
			}()
		}
		for i := 0; i < len(h.KeysList.Keys); i++ {
			if err := <-errC; err != nil {
				return err
			}
		}
	}
	h.WG.Wait()
	close(errC)
	h.IsValidKeys = true
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

func (h *HTTPSessionManager) ViewKeys() {
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

func in(keylist []Key, _key Key) bool {
	for _, key := range keylist {
		if key.ID == _key.ID {
			return true
		}
	}
	return false
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
