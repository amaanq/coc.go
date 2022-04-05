package client

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func New(credentials ...map[string]string) (*HTTPSessionManager, error) {
	var creds []LoginCredential
	for _, credential := range credentials {
		for email, password := range credential {
			creds = append(creds, LoginCredential{Email: email, Password: password})
		}
	}
	H := &HTTPSessionManager{
		client:      resty.New(),
		credentials: creds,
		logins:      make([]LoginResponse, 0),
		keyIndex:    0,
		ready:       true,
	}

	for _, credential := range H.credentials {
		err := H.login(credential)
		if err != nil {
			return nil, err
		}

		err = H.getKeys()
		if err != nil {
			return nil, err
		}

		err = H.updateKeys()
		if err != nil {
			return nil, err
		}
	}

	return H, nil
}

func (h *HTTPSessionManager) login(credential LoginCredential) error {
	resp, err := h.client.R().
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

func (h *HTTPSessionManager) getKeys() error {
	resp, err := h.client.R().
		SetHeader("Content-Type", "application/json").
		Post(DevBaseUrl + KeyListEndpoint)
	if err != nil {
		return err
	}

	if h.StatusCode = resp.StatusCode(); h.StatusCode != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	if err = json.Unmarshal(resp.Body(), &h.currentLoginKeys); err != nil {
		return err
	}

	h.allKeys.Keys = append(h.allKeys.Keys, h.currentLoginKeys.Keys...)
	return nil
}

func (h *HTTPSessionManager) createKey() error {
	if h.iP == "" {
		h.getiP()
	}

	description := fmt.Sprintf("Created on %s", time.Now().Format(time.RFC3339))
	body := fmt.Sprintf(`{"name":"%s","descriPtion":"%s", "cidrRanges": ["%s"], "scopes": ["clash"]}`, "coc.go", description, h.iP)
	resp, err := h.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(DevBaseUrl + KeyCreateEndpoint)
	if err != nil {
		return err
	}

	if h.StatusCode = resp.StatusCode(); h.StatusCode != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	var keycreationresponse KeyResponse
	if err = json.Unmarshal(resp.Body(), &keycreationresponse); err != nil {
		return err
	}

	h.allKeys.Keys = append(h.allKeys.Keys, keycreationresponse.Key)
	return nil
}

func (h *HTTPSessionManager) revokeKey(key Key) error {
	jsonBody := fmt.Sprintf(`{"id": "%s"}`, key.ID)
	resp, err := h.client.R().
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
	removeKey(h.allKeys.Keys, key)

	return nil
}

func (h *HTTPSessionManager) updateKeys() (err error) {
	if h.iP == "" {
		err = h.getiP()
		if err != nil {
			return err
		}
	}

	errs := make(chan error, 10)
	numKeysLeft := 10 - len(h.currentLoginKeys.Keys)

	for i := 0; i < numKeysLeft; i++ {
		h.waitGroup.Add(1)
		go func() {
			defer h.waitGroup.Done()
			err = h.createKey()
			errs <- err
		}()
	}

	h.waitGroup.Wait()

	for _, key := range h.currentLoginKeys.Keys {
		h.waitGroup.Add(1)
		go func(key Key) {
			defer h.waitGroup.Done()
			if !contains(key.Cidrranges, h.iP) {
				err = h.revokeKey(key)
				if err != nil {
					errs <- err
					return
				}
				err = h.createKey()
				if err != nil {
					errs <- err
					return
				}
			}
			errs <- nil
		}(key)
	}

	h.waitGroup.Wait()

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

func (h *HTTPSessionManager) getiP() error {
	resp, err := h.client.R().
		Get(IPUrl)
	if err != nil {
		return err
	}

	if h.StatusCode = resp.StatusCode(); h.StatusCode != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	h.iP = string(resp.Body())
	return nil
}

func removeKey(keylist []Key, deleteKey Key) []Key {
	var ret []Key
	for _, key := range keylist {
		if deleteKey.ID != key.ID {
			ret = append(ret, key)
		}
	}
	return ret
}

func (h *HTTPSessionManager) appendLogin(login LoginResponse) {
	for index, existingLogin := range h.logins {
		if existingLogin.Auth.Uid == login.Auth.Uid { // if this login was already stored just update it, otherwise append it
			h.mutex.Lock()
			h.logins[index] = login
			h.mutex.Unlock()
			return
		}
	}
	h.logins = append(h.logins, login)
}

func (h *HTTPSessionManager) incrementIndex() {
	h.mutex.Lock()
	if h.keyIndex == len(h.allKeys.Keys)-1 {
		h.keyIndex = 0
	} else {
		h.keyIndex += 1
	}
	h.mutex.Unlock()
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
