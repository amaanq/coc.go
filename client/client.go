package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/amaanq/coc.go/clan"
	"github.com/amaanq/coc.go/labels"
	"github.com/amaanq/coc.go/league"
	"github.com/amaanq/coc.go/location"
	"github.com/amaanq/coc.go/player"
)

const (
	BaseUrl = "https://api.clashofclans.com/v1"
)

func (h *HTTPSessionManager) Request(route string, nested bool) ([]byte, error) {
	if !h.ready {
		return nil, fmt.Errorf("keys are not yet ready, wait a few seconds")
	}

	url := BaseUrl + route
	data, contains := h.cache.Get(url)
	if contains {
		byt, ok := data.([]byte)
		if ok {
			return byt, nil
		} else {
			return byt, fmt.Errorf("failed type conversion")
		}
	}

	h.Lock()
	key := h.RawKeysList[h.KeyIndex].Key
	if h.KeyIndex == len(h.RawKeysList)-1 {
		h.KeyIndex = 0
	} else {
		h.KeyIndex += 1
	}
	h.Unlock()

	resp, err := h.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("authorization", fmt.Sprintf("Bearer %s", key)).
		Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 403 {
		if nested {
			return nil, fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
		}
		if strings.Contains(string(resp.Body()), "accessDenied.invalidIp") {
			for _, credential := range h.Credentials {
				err := h.APILopin(credential)
				if err != nil {
					fmt.Println(err.Error())
				}
				err = h.GetKeys()
				if err != nil {
					fmt.Println(err.Error())
				}
				err = h.AddOrDeleteKeysAsNecessary(h.LoginResponse.Developer.ID)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			return h.Request(route, true)
		}
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf(string(resp.Body()))
	}

	cachetime, err := strconv.Atoi(resp.Header().Get("cache-control")[strings.Index(resp.Header().Get("cache-control"), "=")+1:])
	if err != nil {
		fmt.Println(err.Error())
		return resp.Body(), nil
	}
	h.cache.Add(url, resp.Body(), time.Second*time.Duration(cachetime))

	return resp.Body(), nil
}

func (h *HTTPSessionManager) Post(route string, body string, nested bool) ([]byte, error) {
	if !h.ready {
		return nil, fmt.Errorf("keys are not yet ready, wait a few seconds")
	}
	url := BaseUrl + route

	resp, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("authorization", fmt.Sprintf("Bearer %s", h.RawKeysList[h.KeyIndex].Key)).
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == 403 {
		if nested {
			return nil, fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
		}
		if strings.Contains(string(resp.Body()), "accessDenied.invalidIp") {
			for _, credential := range h.Credentials {
				err := h.APILopin(credential)
				if err != nil {
					fmt.Println(err.Error())
				}
				err = h.GetKeys()
				if err != nil {
					fmt.Println(err.Error())
				}
				err = h.AddOrDeleteKeysAsNecessary(h.LoginResponse.Developer.ID)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			return h.Post(route, body, true)
		}
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf(string(resp.Body()))
	}

	if h.KeyIndex == len(h.RawKeysList)-1 {
		h.KeyIndex = 0
	} else {
		h.KeyIndex += 1
	}
	return resp.Body(), nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Clan Methods
//_______________________________________________________________________

func (h *HTTPSessionManager) SearchClans(args ...map[string]string) (*clan.ClanList, ClientError) {
	var clanlist *clan.ClanList
	var err ClientError
	endpoint := "/clans"
	params := parseArgs(args)

	if params == "" {
		err.error = fmt.Errorf("at least one filtering parameter must exist")
		return clanlist, err
	}
	endpoint += params

	data, reqErr := h.Request(endpoint, false)
	if reqErr != nil {
		err.error = reqErr
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, &clanlist); jsonErr != nil {
		err.error = jsonErr
		json.Unmarshal(data, &err)
		return nil, err
	}

	return clanlist, err
}

func (h *HTTPSessionManager) GetClan(ClanTag string) (*clan.Clan, error) {
	ClanTag = CorrectTag(ClanTag)
	var cln *clan.Clan
	endpoint := "/clans/" + url.PathEscape(ClanTag)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &cln); err != nil {
		return nil, err
	}
	return cln, nil
}

func (h *HTTPSessionManager) GetClanMembers(ClanTag string) ([]clan.ClanMember, error) {
	ClanTag = CorrectTag(ClanTag)
	var clanmems clan.ClanMemberEndpoint
	endpoint := "/clans/" + url.PathEscape(ClanTag) + "/members"

	data, err := h.Request(endpoint, false)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &clanmems); err != nil {
		return nil, err
	}
	return clanmems.Items, nil
}

func (h *HTTPSessionManager) GetClanWarLog(ClanTag string) ([]clan.War, error) {
	ClanTag = CorrectTag(ClanTag)
	var clanwarlog *clan.WarLog
	endpoint := "/clans/" + url.PathEscape(ClanTag) + "/warlog"

	data, err := h.Request(endpoint, false)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	if err := json.Unmarshal(data, &clanwarlog); err != nil {
		return nil, err
	}
	return clanwarlog.Items, nil
}

func (h *HTTPSessionManager) GetClanCurrentWar(ClanTag string) (*clan.CurrentWar, error) {
	ClanTag = CorrectTag(ClanTag)
	var clanwar *clan.CurrentWar
	endpoint := "/clans/" + url.PathEscape(ClanTag) + "/currentwar"

	data, err := h.Request(endpoint, false)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	if err := json.Unmarshal(data, &clanwar); err != nil {
		return nil, err
	}
	return clanwar, nil
}

func (h *HTTPSessionManager) GetClanWarLeagueGroup(ClanTag string) { //waiting for next cwl

}

func (h *HTTPSessionManager) GetCWLWars(WarTag string) { //above

}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Player Methods
//_______________________________________________________________________

func (h *HTTPSessionManager) GetPlayer(PlayerTag string) (*player.Player, error) {
	PlayerTag = CorrectTag(PlayerTag)
	var player *player.Player
	endpoint := "/players/" + url.PathEscape(PlayerTag)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return player, err
	}
	if err := json.Unmarshal(data, &player); err != nil {
		return player, err
	}
	return player, nil
}

// Side note: This is the only POST method for the API so far
func (h *HTTPSessionManager) VerifyPlayerToken(PlayerTag string, Token string) (bool, error) {
	PlayerTag = CorrectTag(PlayerTag)
	var verification *player.Verification
	endpoint := "/players/" + url.PathEscape(PlayerTag) + "/verifytoken"

	data, err := h.Post(endpoint, fmt.Sprintf(`{"token": "%s"}`, Token), false)
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(data, &verification); err != nil {
		return false, err
	}
	return verification.Status == "ok", nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// League Methods
//_______________________________________________________________________

func (h *HTTPSessionManager) GetLeagues(args ...map[string]string) (*league.LeagueData, error) {
	var leaguedata league.LeagueData
	endpoint := "/leagues" + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return &leaguedata, err
	}
	if err := json.Unmarshal(data, &leaguedata); err != nil {
		return &leaguedata, err
	}
	return &leaguedata, nil
}

func (h *HTTPSessionManager) GetLeague(LeagueID string) (*league.League, error) {
	var league league.League
	endpoint := "/leagues/" + LeagueID

	data, err := h.Request(endpoint, false)
	if err != nil {
		return &league, err
	}
	if err := json.Unmarshal(data, &league); err != nil {
		return &league, err
	}
	return &league, nil
}

func (h *HTTPSessionManager) GetLeagueSeasons(LeagueID string, args ...map[string]string) (*league.SeasonData, error) {
	var seasondata *league.SeasonData
	if LeagueID != "29000022" {
		fmt.Println("Only Legends League is supported with this. Deferring to 29000022. To avoid this message being printed again, pass in 29000022 for the LeagueID argument.")
		LeagueID = "29000022"
	}
	endpoint := "/leagues/" + LeagueID + "/seasons" + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return seasondata, err
	}
	if err := json.Unmarshal(data, &seasondata); err != nil {
		return seasondata, err
	}
	return seasondata, nil
}

// Be cautious when using this. the data returned is massive. recommended to add args of {"limit": limit} and use the cursors for more data
func (h *HTTPSessionManager) GetLeagueSeasonInfo(LeagueID string, SeasonID string, args ...map[string]string) (*league.SeasonInfo, error) {
	var seasoninfo *league.SeasonInfo
	if LeagueID != "29000022" {
		fmt.Println("Only Legends League is supported with this. Deferring to 29000022. To avoid this message being printed again, pass in 29000022 for the LeagueID argument.")
		LeagueID = "29000022"
	}
	match, err := regexp.MatchString("^20[0-2][0-9]-((0[1-9])|(1[0-2]))$", SeasonID)
	if err != nil {
		return seasoninfo, err
	}
	if !match {
		return seasoninfo, fmt.Errorf("invalid season format, format must match YYYY-MM")
	}
	endpoint := "/leagues/" + LeagueID + "/seasons/" + SeasonID + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return seasoninfo, err
	}
	if err := json.Unmarshal(data, &seasoninfo); err != nil {
		return seasoninfo, err
	}
	return seasoninfo, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Location Methods
//_______________________________________________________________________

//This should be passed ideally with nothing, kwargs aren't necessary here but only for the sake of completeness.
func (h *HTTPSessionManager) GetLocations(args ...map[string]string) (*location.LocationData, error) {
	var locationdata *location.LocationData
	endpoint := "/locations" + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return locationdata, err
	}
	if err := json.Unmarshal(data, &locationdata); err != nil {
		return locationdata, err
	}
	return locationdata, nil
}

func (h *HTTPSessionManager) GetLocation(LocationID string) (*location.Location, error) {
	var location *location.Location
	endpoint := "/locations/" + LocationID

	data, err := h.Request(endpoint, false)
	if err != nil {
		return location, err
	}
	if err := json.Unmarshal(data, &location); err != nil {
		return location, err
	}
	return location, nil
}

// Main Village Clan Rankings
func (h *HTTPSessionManager) GetLocationClans(LocationID string, args ...map[string]string) (*location.ClanData, error) {
	var clandata *location.ClanData
	endpoint := "/locations/" + LocationID + "/rankings/clans" + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return clandata, err
	}
	if err := json.Unmarshal(data, &clandata); err != nil {
		return clandata, err
	}
	return clandata, nil
}

// Builder Hall Clan Rankings
func (h *HTTPSessionManager) GetLocationClansVersus(LocationID string, args ...map[string]string) (*location.ClanVersusData, error) {
	var clanversusdata *location.ClanVersusData
	endpoint := "/locations/" + LocationID + "/rankings/clans-versus" + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return clanversusdata, err
	}
	if err := json.Unmarshal(data, &clanversusdata); err != nil {
		return clanversusdata, err
	}
	return clanversusdata, nil
}

// Main Village Player Rankings
func (h *HTTPSessionManager) GetLocationPlayers(LocationID string, args ...map[string]string) (*location.PlayerData, error) {
	var playerdata *location.PlayerData
	endpoint := "/locations/" + LocationID + "/rankings/players" + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return playerdata, err
	}
	if err := json.Unmarshal(data, &playerdata); err != nil {
		return playerdata, err
	}
	return playerdata, nil
}

// Builder Hall Player Rankings
func (h *HTTPSessionManager) GetLocationPlayersVersus(LocationID string, args ...map[string]string) (*location.PlayerVersusData, error) {
	var playerversusdata *location.PlayerVersusData
	endpoint := "/locations/" + LocationID + "/rankings/players-versus" + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return playerversusdata, err
	}
	if err := json.Unmarshal(data, &playerversusdata); err != nil {
		return playerversusdata, err
	}
	return playerversusdata, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Label Methods
//_______________________________________________________________________

func (h *HTTPSessionManager) GetClanLabels(args ...map[string]string) (*labels.LabelsData, error) {
	var labels *labels.LabelsData
	endpoint := "/labels/clans" + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return labels, err
	}
	if err := json.Unmarshal(data, &labels); err != nil {
		return labels, err
	}
	return labels, nil
}

func (h *HTTPSessionManager) GetPlayerLabels(args ...map[string]string) (*labels.LabelsData, error) {
	var labels *labels.LabelsData
	endpoint := "/labels/players" + parseArgs(args)

	data, err := h.Request(endpoint, false)
	if err != nil {
		return labels, err
	}
	if err := json.Unmarshal(data, &labels); err != nil {
		return labels, err
	}
	return labels, nil
}
