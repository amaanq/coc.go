package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/amaanq/coc.go/clan"
	"github.com/amaanq/coc.go/labels"
	"github.com/amaanq/coc.go/league"
	"github.com/amaanq/coc.go/location"
	"github.com/amaanq/coc.go/player"
)

func (h *HTTPSessionManager) request(route string, nested bool) ([]byte, error) {
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

	h.incrementIndex()
	key := h.allKeys.Keys[h.keyIndex].Key

	resp, err := h.client.R().
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
			for _, credential := range h.credentials {
				err := h.login(credential)
				if err != nil {
					return nil, err
				}

				err = h.getKeys()
				if err != nil {
					return nil, err
				}

				err = h.updateKeys()
				if err != nil {
					return nil, err
				}
			}
			return h.request(route, true)
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

func (h *HTTPSessionManager) post(route string, body string, nested bool) ([]byte, error) {
	if !h.ready {
		return nil, fmt.Errorf("keys are not yet ready, wait a few seconds")
	}
	url := BaseUrl + route

	h.incrementIndex()
	key := h.allKeys.Keys[h.keyIndex].Key

	resp, err := h.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("authorization", fmt.Sprintf("Bearer %s", key)).
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
			for _, credential := range h.credentials {
				err := h.login(credential)
				if err != nil {
					return nil, err
				}

				err = h.getKeys()
				if err != nil {
					return nil, err
				}

				err = h.updateKeys()
				if err != nil {
					return nil, err
				}
			}
			return h.post(route, body, true)
		}
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf(string(resp.Body()))
	}

	return resp.Body(), nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Clan Methods
//_______________________________________________________________________

func (h *HTTPSessionManager) SearchClans(args ...map[string]string) (ClanList *clan.ClanList, err ClientError) {
	ClanList = &clan.ClanList{}
	params := parseArgs(args)
	if params == "" {
		err.setErr(fmt.Errorf("at least one filtering parameter must exist"))
		return nil, err
	}

	data, reqErr := h.request(ClanEndpoint+params, false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, ClanList); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

func (h *HTTPSessionManager) GetClan(ClanTag string) (Clan *clan.Clan, err ClientError) {
	Clan = &clan.Clan{}
	ClanTag = CorrectTag(ClanTag)

	data, reqErr := h.request(ClanEndpoint+"/"+url.PathEscape(ClanTag), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, Clan); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

func (h *HTTPSessionManager) GetClanMembers(ClanTag string) (ClanMems []clan.ClanMember, err ClientError) {
	ClanMems = []clan.ClanMember{}
	ClanTag = CorrectTag(ClanTag)

	data, reqErr := h.request(ClanEndpoint+"/"+url.PathEscape(ClanTag)+"/members/", false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, &ClanMems); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

func (h *HTTPSessionManager) GetClanWarLog(ClanTag string) (ClanWarLog *clan.WarLog, err ClientError) {
	ClanWarLog = &clan.WarLog{}
	ClanTag = CorrectTag(ClanTag)

	data, reqErr := h.request(ClanEndpoint+"/"+url.PathEscape(ClanTag)+"/warlog/", false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, ClanWarLog); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

func (h *HTTPSessionManager) GetClanCurrentWar(ClanTag string) (ClanWar *clan.CurrentWar, err ClientError) {
	ClanWar = &clan.CurrentWar{}
	ClanTag = CorrectTag(ClanTag)

	data, reqErr := h.request(ClanEndpoint+"/"+url.PathEscape(ClanTag)+"/currentwar/", false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, ClanWar); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

func (h *HTTPSessionManager) GetClanWarLeagueGroup(ClanTag string) { //waiting for next cwl

}

func (h *HTTPSessionManager) GetCWLWars(WarTag string) { //above

}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Player Methods
//_______________________________________________________________________

func (h *HTTPSessionManager) GetPlayer(PlayerTag string) (Player *player.Player, err ClientError) {
	Player = &player.Player{}
	PlayerTag = CorrectTag(PlayerTag)

	data, reqErr := h.request(PlayerEndpoint+"/"+url.PathEscape(PlayerTag), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, Player); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

// A slice of players will be returned with no error, however players that had an error (i.e are banned or the tag never existed)
// will be nil inside the slice. The order of the player tags is kept intact.
func (h *HTTPSessionManager) GetPlayers(PlayerTags []string) []*player.Player {
	Players := make([]*player.Player, len(PlayerTags))
	PlayerMap := make(map[string]*player.Player)

	for idx, tag := range PlayerTags { // correct all tags to check with the map later
		PlayerTags[idx] = CorrectTag(tag)
	}

	var playerWg sync.WaitGroup
	var playerMapMutex sync.Mutex

	playerWg.Add(len(PlayerTags))
	for _, tag := range PlayerTags {
		go func(t string) {
			defer playerWg.Done()
			player, err := h.GetPlayer(t)
			if err.Err() != nil {
				playerMapMutex.Lock()
				PlayerMap[t] = nil
				playerMapMutex.Unlock()
				return
			}
			playerMapMutex.Lock()
			PlayerMap[t] = player
			playerMapMutex.Unlock()
		}(tag)
	}
	playerWg.Wait()

	for i, tag := range PlayerTags {
		Players[i] = PlayerMap[tag]
	}
	return Players
}

// Side note: This is the only POST method for the API so far
func (h *HTTPSessionManager) VerifyPlayerToken(PlayerTag string, Token string) (Verification *player.Verification, err ClientError) {
	Verification = &player.Verification{}
	PlayerTag = CorrectTag(PlayerTag)

	data, reqErr := h.post(PlayerEndpoint+"/"+url.PathEscape(PlayerTag)+"/verifytoken/", fmt.Sprintf(`{"token": "%s"}`, Token), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, Verification); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// League Methods
//_______________________________________________________________________

func (h *HTTPSessionManager) GetLeagues(args ...map[string]string) (LeagueData *league.LeagueData, err ClientError) {
	LeagueData = &league.LeagueData{}
	data, reqErr := h.request(LeagueEndpoint+parseArgs(args), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, LeagueData); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

func (h *HTTPSessionManager) GetLeague(LeagueID string) (League *league.League, err ClientError) {
	League = &league.League{}
	data, reqErr := h.request(LeagueEndpoint+"/"+LeagueID, false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, League); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

// Ensure you pass in 29000022 for LeagueID
func (h *HTTPSessionManager) GetLeagueSeasons(LeagueID league.LeagueID, args ...map[string]string) (SeasonData *league.SeasonData, err ClientError) {
	SeasonData = &league.SeasonData{}
	if LeagueID != league.LegendLeague {
		fmt.Println("Only Legends League is supported with this. Deferring to 29000022. To avoid this message being printed again, pass in 29000022 for the LeagueID argument.")
		LeagueID = league.LegendLeague
	}

	data, reqErr := h.request(LeagueEndpoint+"/"+fmt.Sprint(LeagueID)+"/seasons"+parseArgs(args), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, SeasonData); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

// Be cautious when using this. the data returned is massive. recommended to add args of {"limit": limit} and use the cursors for more data
func (h *HTTPSessionManager) GetLeagueSeasonInfo(LeagueID league.LeagueID, SeasonID string, args ...map[string]string) (SeasonInfo *league.SeasonInfo, err ClientError) {
	SeasonInfo = &league.SeasonInfo{}
	if LeagueID != league.LegendLeague {
		fmt.Println("Only Legends League is supported with this. Deferring to 29000022. To avoid this message being printed again, pass in 29000022 for the LeagueID argument.")
		LeagueID = league.LegendLeague
	}

	match, matchErr := regexp.MatchString("^20[0-2][0-9]-((0[1-9])|(1[0-2]))$", SeasonID)
	if matchErr != nil {
		err.setErr(matchErr)
		return nil, err
	}

	if !match {
		err.setErr(fmt.Errorf("invalid season format, format must match the YYYY-MM date format"))
		return nil, err
	}

	data, reqErr := h.request(LeagueEndpoint+"/"+fmt.Sprint(LeagueID)+"/seasons/"+SeasonID+parseArgs(args), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, SeasonInfo); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Location Methods
//_______________________________________________________________________

//This should be passed ideally with nothing, kwargs aren't necessary here but only for the sake of completeness.
func (h *HTTPSessionManager) GetLocations(args ...map[string]string) (LocationData *location.LocationData, err ClientError) {
	LocationData = &location.LocationData{}
	data, reqErr := h.request(LocationEndpoint+parseArgs(args), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, LocationData); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

func (h *HTTPSessionManager) GetLocation(LocationID location.LocationID) (Location *location.Location, err ClientError) {
	Location = &location.Location{}
	data, reqErr := h.request(LocationEndpoint+"/"+fmt.Sprint(LocationID), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, Location); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

// Main Village Clan Rankings
func (h *HTTPSessionManager) GetLocationClans(LocationID location.LocationID, args ...map[string]string) (ClanData *location.ClanData, err ClientError) {
	ClanData = &location.ClanData{}
	data, reqErr := h.request(LocationEndpoint+"/"+fmt.Sprint(LocationID)+"/rankings/clans"+parseArgs(args), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, ClanData); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

// Builder Hall Clan Rankings
func (h *HTTPSessionManager) GetLocationClansVersus(LocationID location.LocationID, args ...map[string]string) (ClanVersusData *location.ClanVersusData, err ClientError) {
	ClanVersusData = &location.ClanVersusData{}
	data, reqErr := h.request(LocationEndpoint+"/"+fmt.Sprint(LocationID)+"/rankings/clans-versus"+parseArgs(args), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, ClanVersusData); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

// Main Village Player Rankings
func (h *HTTPSessionManager) GetLocationPlayers(LocationID location.LocationID, args ...map[string]string) (PlayerData *location.PlayerData, err ClientError) {
	PlayerData = &location.PlayerData{}
	data, reqErr := h.request(LocationEndpoint+"/"+fmt.Sprint(LocationID)+"/rankings/players"+parseArgs(args), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, PlayerData); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

// Builder Hall Player Rankings
func (h *HTTPSessionManager) GetLocationPlayersVersus(LocationID location.LocationID, args ...map[string]string) (PlayerVersusData *location.PlayerVersusData, err ClientError) {
	PlayerVersusData = &location.PlayerVersusData{}
	data, reqErr := h.request(LocationEndpoint+fmt.Sprint(LocationID)+"/rankings/players-versus", false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, PlayerVersusData); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Label Methods
//_______________________________________________________________________

func (h *HTTPSessionManager) GetClanLabels(args ...map[string]string) (LabelsData *labels.LabelsData, err ClientError) {
	LabelsData = &labels.LabelsData{}
	data, reqErr := h.request(LabelEndpoint+"/clans"+parseArgs(args), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, LabelsData); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}

func (h *HTTPSessionManager) GetPlayerLabels(args ...map[string]string) (LabelsData *labels.LabelsData, err ClientError) {
	LabelsData = &labels.LabelsData{}
	data, reqErr := h.request(LabelEndpoint+"/players"+parseArgs(args), false)
	if reqErr != nil {
		err.setErr(reqErr)
		return nil, err
	}

	if jsonErr := json.Unmarshal(data, LabelsData); jsonErr != nil {
		err.setErr(jsonErr)
		json.Unmarshal(data, &err)
		return nil, err
	}
	return
}
