package coc

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

func newClient(credentials map[string]string) (*Client, error) {
	var accounts []APIAccount
	for email, password := range credentials {
		// For each email, create a login where the actual credentials are filled out, to be fully filled in later when logging in and getting keys.
		accounts = append(accounts, APIAccount{
			Credential: Credential{Email: email, Password: password},
		})
	}

	H := &Client{
		client:     resty.New(),
		ready:      true,
		accounts:   accounts,
		index:      Index{KeyAccountIndex: 0, KeyIndex: 0},
		StatusCode: 0,
		ipAddress:  "",
	}
	H.getIP()

	for index := range H.accounts {
		err := H.accounts[index].login(H.client)
		if err != nil {
			return nil, err
		}

		// This calls getKeys() anyways
		err = H.accounts[index].updateKeys(H.ipAddress, H.client)
		if err != nil {
			return nil, err
		}
	}

	return H, nil
}

// Log in to an account and unmarshal response into a.Response
func (a *APIAccount) login(client *resty.Client) error {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"email":"%s","password":"%s"}`, a.Credential.Email, a.Credential.Password)).
		Post(DevBaseUrl + LoginEndpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	if err = json.Unmarshal(resp.Body(), &a.Response); err != nil {
		return err
	}

	return nil
}

func (a *APIAccount) getKeys(client *resty.Client) error {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Post(DevBaseUrl + KeyListEndpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	if err = json.Unmarshal(resp.Body(), &a.Keys); err != nil {
		return err
	}

	return nil
}

func (a *APIAccount) createKey(ip string, client *resty.Client) error {
	description := fmt.Sprintf("Created on %s", time.Now().Format(time.RFC3339))
	body := fmt.Sprintf(`{"name":"%s","descriPtion":"%s", "cidrRanges": ["%s"], "scopes": ["clash"]}`, "coc.go", description, ip)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(DevBaseUrl + KeyCreateEndpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	var keycreationresponse KeyResponse
	if err = json.Unmarshal(resp.Body(), &keycreationresponse); err != nil {
		return err
	}

	// Refresh account's keys after creation here.
	go a.getKeys(client)

	return nil
}

func (a *APIAccount) revokeKey(key Key, client *resty.Client) error {
	jsonBody := fmt.Sprintf(`{"id": "%s"}`, key.ID)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonBody).
		Post(DevBaseUrl + KeyRevokeEndpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	var keydeletionresponse KeyResponse
	if err = json.Unmarshal(resp.Body(), &keydeletionresponse); err != nil {
		return err
	}

	// Refresh account's keys after deletion here.
	go a.getKeys(client)

	return nil
}

// This will create a key if there aren't 10 keys already for this login.
// It will also revoke and recreate keys if the IP address changes.
func (a *APIAccount) updateKeys(ip string, client *resty.Client) error {
	err := a.getKeys(client) // check keys first
	if err != nil {
		return err
	}

	errs := make(chan error, 10)
	numKeysLeft := 10 - len(a.Keys.Keys)
	var wg sync.WaitGroup

	// First re-create keys if the IP address changes.
	for _, key := range a.Keys.Keys {
		if !contains(key.Cidrranges, ip) {
			wg.Add(1)
			go func(key Key) {
				defer wg.Done()
				err := a.revokeKey(key, client)
				if err != nil {
					errs <- err
					return
				}
				err = a.createKey(ip, client)
				if err != nil {
					errs <- err
					return
				}
			}(key)
		}
	}
	wg.Wait()

	// Then create more keys if there aren't 10.
	wg.Add(numKeysLeft)
	for i := 0; i < numKeysLeft; i++ {
		go func() {
			defer wg.Done()
			err := a.createKey(ip, client)
			if err != nil {
				errs <- err
			}
		}()
	}
	wg.Wait()

	if len(errs) > 0 {
		return <-errs
	}
	return nil
}

func (c *Client) getIP() error {
	resp, err := c.client.R().
		Get(IPUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	c.ipAddress = string(resp.Body())
	return nil
}

func (c *Client) incIndex() {
	c.Lock()
	if c.index.KeyIndex == len(c.accounts[c.index.KeyAccountIndex].Keys.Keys)-1 {
		c.index.KeyIndex = 0
		c.index.KeyAccountIndex++
	} else {
		c.index.KeyIndex++
	}
	c.Unlock()
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
