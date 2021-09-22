package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/amaanq/coc.go/clan"
	"github.com/amaanq/coc.go/location"
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

func (h *HTTPSessionManager) SearchClans(arg map[string]string) ([]byte, error) {
	endpoint := "/clans"
	params := ""
	for key, val := range arg {
		params += fmt.Sprintf("%s=%s&", key, val)
	}
	params = params[:len(params)-1]
	if params != "" {
		endpoint += "?" + params
	}
	data, err := h.Request(endpoint, false)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (h *HTTPSessionManager) GetClan(ClanTag string) (clan.Clan, error) {
	if !strings.Contains(ClanTag, "#") {
		ClanTag = "#" + ClanTag
	}
	var cln clan.Clan
	endpoint := "/clans/" + url.PathEscape(ClanTag)
	data, err := h.Request(endpoint, false)
	if err != nil {
		return cln, err
	}
	if err := json.Unmarshal(data, &cln); err != nil {
		return cln, err
	}
	return cln, nil
}

func (h *HTTPSessionManager) GetClanMembers(ClanTag string) (clan.MemberList, error) {
	if !strings.Contains(ClanTag, "#") {
		ClanTag = "#" + ClanTag
	}
	var clanmems clan.MemberList
	endpoint := "/clans/" + url.PathEscape(ClanTag) + "/members"
	data, err := h.Request(endpoint, false)
	if err != nil {
		return clanmems, err
	}
	if err := json.Unmarshal(data, &clanmems); err != nil {
		return clanmems, err
	}
	return clanmems, nil
}

func (h *HTTPSessionManager) GetClanWarLog(ClanTag string) (clan.WarLog, error) {
	if !strings.Contains(ClanTag, "#") {
		ClanTag = "#" + ClanTag
	}
	var clanwarlog clan.WarLog
	endpoint := "/clans/" + url.PathEscape(ClanTag) + "/warlog"
	data, err := h.Request(endpoint, false)
	if err != nil {
		return clanwarlog, err
	}
	if err := json.Unmarshal(data, &clanwarlog); err != nil {
		return clanwarlog, err
	}
	return clanwarlog, nil
}

func (h *HTTPSessionManager) GetClanCurrentWar(ClanTag string) (clan.CurrentWar, error) {
	if !strings.Contains(ClanTag, "#") {
		ClanTag = "#" + ClanTag
	}
	var clanwar clan.CurrentWar
	endpoint := "/clans/" + url.PathEscape(ClanTag) + "/currentwar"
	data, err := h.Request(endpoint, false)
	if err != nil {
		return clanwar, err
	}
	if err := json.Unmarshal(data, &clanwar); err != nil {
		return clanwar, err
	}
	return clanwar, nil
}

func (h *HTTPSessionManager) GetClanWarLeagueGroup(ClanTag string) { //waiting for next cwl

}

func (h *HTTPSessionManager) GetCWLWars(WarTag string) { //above

}

//This should be passed ideally with nothing, kwargs aren't necessary here but only for the sake of completeness.
func (h *HTTPSessionManager) SearchLocations(args ...map[string]string) (location.LocationData, error) {
	var locationdata location.LocationData
	endpoint := "/locations"+parseArgs(args)
	data, err := h.Request(endpoint, false)
	if err != nil {
		return locationdata, err
	}
	if err := json.Unmarshal(data, &locationdata); err != nil {
		return locationdata, err
	}
	return locationdata, nil
}

func (h *HTTPSessionManager) GetLocation(LocationID string) (location.Location, error) {
	var location location.Location
	endpoint := "/locations/"+LocationID
	data, err := h.Request(endpoint, false)
	if err != nil {
		return location, err
	}
	if err := json.Unmarshal(data, &location); err != nil {
		return location, err
	}
	return location, nil
}

func (h *HTTPSessionManager) GetLocationClans(LocationID string, args ...map[string]string) (location.ClanData, error){
	var clandata location.ClanData
	endpoint := "/locations/"+LocationID+"/rankings/clans"+parseArgs(args)
	data, err := h.Request(endpoint, false)
	if err != nil {
		return clandata, err
	}
	if err := json.Unmarshal(data, &clandata); err != nil {
		return clandata, err
	}
	return clandata, nil
}

func (h *HTTPSessionManager) GetLocationPlayers(LocationID string, args ...map[string]string) (location.PlayerData, error) {
	var playerdata location.PlayerData
	endpoint := "/locations/"+LocationID+"/rankings/players"+parseArgs(args)
	data, err := h.Request(endpoint, false)
	if err != nil {
		return playerdata, err
	}
	if err := json.Unmarshal(data, &playerdata); err != nil {
		return playerdata, err
	}
	return playerdata, nil
}

func (h *HTTPSessionManager) GetLocationClansVersus(LocationID string, args ...map[string]string) (location.ClanVersusData, error) {
	var clanversusdata location.ClanVersusData
	endpoint := "/locations/"+LocationID+"/rankings/clans-versus"+parseArgs(args)
	data, err := h.Request(endpoint, false)
	if err != nil {
		return clanversusdata, err
	}
	if err := json.Unmarshal(data, &clanversusdata); err != nil {
		return clanversusdata, err
	}
	return clanversusdata, nil
}

func (h *HTTPSessionManager) GetLocationPlayersVersus(LocationID string, args ...map[string]string) (location.PlayerVersusData, error) {
	var playerversusdata location.PlayerVersusData
	endpoint := "/locations/"+LocationID+"/rankings/players-versus"+parseArgs(args)
	data, err := h.Request(endpoint, false)
	if err != nil {
		return playerversusdata, err
	}
	if err := json.Unmarshal(data, &playerversusdata); err != nil {
		return playerversusdata, err
	}
	return playerversusdata, nil
}

func (h *HTTPSessionManager) SearchLeagues(args ...map[string]string) {

}

func (h *HTTPSessionManager) GetLeague(LeagueID string) {

}

func (h *HTTPSessionManager) GetLeagueSeasons(LeagueID string, args ...map[string]string) {

}

func (h *HTTPSessionManager) GetLeagueSeasonInfo(LeagueID string, SeasonID string, args ...map[string]string) {

}

func (h *HTTPSessionManager) GetPlayer(PlayerTag string) {

}

func (h *HTTPSessionManager) VerifyPlayerToken(PlayerTag string, Token string) {

}

func (h *HTTPSessionManager) GetClanLabels(args ...map[string]string) {

}

func (h *HTTPSessionManager) GetPlayerLabels(args ...map[string]string) {

}

func parseArgs(args []map[string]string) string {
	if len(args) == 0 {
		return ""
	}
	params := ""
	if len(args) > 0 {
		for _, arg := range args {
			for key, val := range arg {
				params += fmt.Sprintf("%s=%s&", key, val)
			}
		}
	}
	params = params[:len(params)-1]
	if params != "" {
		params = "?" + params
	}
	return params
}
