package client

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
)

func New(credentials ...map[string]string) (*HTTPSessionManager, error) {
	var creds []LoginCredential
	for _, credential := range credentials {
		for email, password := range credential {
			creds = append(creds, LoginCredential{Email: email, Password: password})
		}
	}
	H := &HTTPSessionManager{
		Credentials: creds,
		Logins:      make([]LoginResponse, 0),
		KeyIndex:    0,
		ready:       true,
		cache:       cache.New(time.Second*60, time.Second*60),
	}
	H.Client = resty.New()

	for _, credential := range H.Credentials {
		err := H.Login(credential)
		if err != nil {
			return nil, err
		}

		err = H.GetKeys()
		if err != nil {
			return nil, err
		}

		err = H.UpdateKeys()
		if err != nil {
			return nil, err
		}
	}

	return H, nil
}

func (h *HTTPSessionManager) Login(credential LoginCredential) error {
	resp, err := h.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"email":"%s","password":"%s"}`, credential.Email, credential.Password)).
		Post(DevBaseUrl + LoginEndpoint)
	if err != nil {
		return err
	}

	if h.StatusCode = resp.StatusCode(); h.StatusCode != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	var loginresponse LoginResponse
	if err = json.Unmarshal(resp.Body(), &loginresponse); err != nil {
		return err
	}

	h.appendLogin(loginresponse)

	return nil
}

func (h *HTTPSessionManager) GetKeys() error {
	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		Post(DevBaseUrl + KeyListEndpoint)
	if err != nil {
		return err
	}

	if h.StatusCode = resp.StatusCode(); h.StatusCode != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	if err = json.Unmarshal(resp.Body(), &h.CurrentLoginKeys); err != nil {
		return err
	}

	h.AllKeys.Keys = append(h.AllKeys.Keys, h.CurrentLoginKeys.Keys...)

	fmt.Println("KEYLIST: ", string(resp.Body()))
	return nil
}

func (h *HTTPSessionManager) CreateKey() error {
	if h.IP == "" {
		h.getIP()
	}

	description := fmt.Sprintf("Created on %s", time.Now().Format(time.RFC3339))
	body := fmt.Sprintf(`{"name":"%s","description":"%s", "cidrRanges": ["%s"], "scopes": ["clash"]}`, "coc.go", description, h.IP)
	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(DevBaseUrl + KeyCreateEndpoint)
	if err != nil {
		return err
	}

	fmt.Println(DevBaseUrl + KeyCreateEndpoint)

	if h.StatusCode = resp.StatusCode(); h.StatusCode != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	var keycreationresponse KeyResponse
	if err = json.Unmarshal(resp.Body(), &keycreationresponse); err != nil {
		return err
	}

	h.AllKeys.Keys = append(h.AllKeys.Keys, keycreationresponse.Key)

	fmt.Println("Added key with ID", keycreationresponse.Key.ID)
	return nil
}

func (h *HTTPSessionManager) RevokeKey(key Key) error {
	jsonBody := fmt.Sprintf(`{"id": "%s"}`, key.ID)
	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonBody).
		Post(DevBaseUrl + KeyRevokeEndpoint)
	if err != nil {
		return err
	}

	if h.StatusCode = resp.StatusCode(); h.StatusCode != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	var keydeletionresponse KeyResponse
	if err = json.Unmarshal(resp.Body(), &keydeletionresponse); err != nil {
		return err
	}

	return nil
}

func (h *HTTPSessionManager) UpdateKeys() (err error) {
	if h.IP == "" {
		err = h.getIP()
		if err != nil {
			return err
		}
	}

	errs := make(chan error, 10)
	var keyWG sync.WaitGroup

	numKeysLeft := 10 - len(h.CurrentLoginKeys.Keys)

	for i := 0; i < numKeysLeft; i++ {
		keyWG.Add(1)
		go func() {
			defer keyWG.Done()
			err = h.CreateKey()
			errs <- err
		}()
	}

	keyWG.Wait()

	for _, key := range h.CurrentLoginKeys.Keys {
		keyWG.Add(1)
		go func(key Key) {
			defer keyWG.Done()
			if !contains(key.Cidrranges, h.IP) {
				err = h.RevokeKey(key)
				if err != nil {
					errs <- err
					return
				}
				err = h.CreateKey()
				if err != nil {
					errs <- err
					return
				}
			}
			errs <- nil
		}(key)
	}

	keyWG.Wait()

	for {
		select {
		case e := <-errs:
			close(errs)
			return e
		case <-time.After(time.Millisecond * 500):
			close(errs)
			return nil
		}
	}
}

func (h *HTTPSessionManager) getIP() error {
	resp, err := h.Client.R().
		Get(IPUrl)
	if err != nil {
		return err
	}

	if h.StatusCode = resp.StatusCode(); h.StatusCode != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	h.IP = string(resp.Body())
	return nil
}

func (h *HTTPSessionManager) ViewKeys() {
	for _, key := range h.AllKeys.Keys {
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

func (h *HTTPSessionManager) appendLogin(login LoginResponse) {
	for index, existingLogin := range h.Logins {
		if existingLogin.Auth.Uid == login.Auth.Uid { // if this login was already stored just update it, otherwise append it
			h.Lock()
			h.Logins[index] = login
			h.Unlock()
			return
		}
	}
	h.Logins = append(h.Logins, login)
}

func (h *HTTPSessionManager) incrementIndex() {
	h.Lock()
	if h.KeyIndex == len(h.AllKeys.Keys)-1 {
		h.KeyIndex = 0
	} else {
		h.KeyIndex += 1
	}
	h.Unlock()
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
