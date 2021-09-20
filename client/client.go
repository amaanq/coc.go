package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/amaanq/coc.go/clan"
	"github.com/go-resty/resty/v2"
)

const (
	BaseUrl = "https://api.clashofclans.com/v1"
)

func (h *HTTPSessionManager) Request(route string, nested bool) ([]byte, error) {
	url := BaseUrl + route
	var req *resty.Request
	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("authorization", fmt.Sprintf("Bearer %s", h.KeysList.Keys[h.KeyIndex].Key)).
		SetResult(&req).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == 403 {
		if nested {
			return nil, fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
		}
		if strings.Contains(string(resp.Body()), "accessDenied.invalidIp") {
			err := h.AddOrDeleteKeysAsNecessary()
			if err != nil {
				return nil, err 
			}
			return h.Request(route, true)
		} 
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
	}
	return resp.Body(), nil
}

func (h *HTTPSessionManager) SearchClans(name string) ([]byte, error) {
	data, err := h.Request("/clans", false)
	if err != nil {
		return nil, err
	}
}

func (h *HTTPSessionManager) GetClan(name string) (clan.Clan, error) {
	var cln clan.Clan
	data, err := h.Request("/clans", false)
	if err != nil {
		return cln, err 
	}
	if err := json.Unmarshal(data, &cln); err != nil {
		return cln, err 
	}
	return cln, nil 
}

func (h *HTTPSessionManager) GetClanMembers(tag string) ([]byte, error) {
	if !strings.Contains(tag, "#") {
		tag = "#"+tag
	}
	tag := url.PathEscape(tag)
	endpoint := "/clans/"+tag+"/members"
	return h.Request(endpoint, false)
}

func (h *HTTPSessionManager) GetClan(name string) ([]byte, error) {
	return h.Request("/clans", false)
}

func (h *HTTPSessionManager) GetClan(name string) ([]byte, error) {
	return h.Request("/clans", false)
}