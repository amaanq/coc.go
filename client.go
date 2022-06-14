package coc

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

const (
	GET  = "GET"
	POST = "POST"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (h *Client) do(method, route, body string, nested bool) ([]byte, error) {
	if !h.ready {
		return nil, fmt.Errorf("keys are not yet ready, wait a few seconds")
	}

	url := BaseUrl + route

	if useCache {
		if data, err := getFromCache(url); err == nil {
			return data, nil
		}
	}

	key := h.accounts[h.index.KeyAccountIndex].Keys.Keys[h.index.KeyIndex].Key
	h.incIndex()

	req := h.client.R().SetHeaders(map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"authorization": fmt.Sprintf("Bearer %s", key),
	})

	var resp *resty.Response
	var err error
	switch method {
	case "GET":
		resp, err = req.Get(url)
	case "POST":
		req.SetBody(body)
		resp, err = req.Post(url)
	}

	if err != nil {
		return nil, err
	}

	var APIError APIError
	if resp.StatusCode() >= 300 { // bind error anyways beforehand...
		if err := json.Unmarshal(resp.Body(), &APIError); err != nil {
			return nil, err
		}
	}

	if resp.StatusCode() == 403 {
		if nested {
			return nil, &APIError //fmt.Errorf(fmt.Sprintf("[%d]: %s", resp.StatusCode(), string(resp.Body())))
		}

		if APIError.Reason == InvalidIP {
			h.getIP()
			for index := range h.accounts {
				err := h.accounts[index].login(h.client)
				if err != nil {
					return nil, err
				}

				// This calls getKeys() anyways
				err = h.accounts[index].updateKeys(h.ipAddress, h.client)
				if err != nil {
					return nil, err
				}
			}
			return h.do(method, route, body, true) // Nested = true so return on error no matter what here (we don't want to loop)
		}
	}

	if resp.StatusCode() >= 300 {
		return nil, &APIError
	}

	cachetime, err := strconv.Atoi(resp.Header().Get("cache-control")[strings.Index(resp.Header().Get("cache-control"), "=")+1:])
	if err != nil {
		fmt.Println(err.Error())
		return resp.Body(), nil
	}

	if useCache {
		if err = writeToCache(url, resp.Body(), cachetime); err != nil {
			return nil, err
		}
	}
	return resp.Body(), nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Clan Methods
//_______________________________________________________________________

func (h *Client) SearchClans(options *clanSearchOptions) (*ClanList, error) {
	var ClanList ClanList
	var opts string

	if options != nil {
		opts = options.ToQuery()
	}
	if opts == "" {
		return nil, fmt.Errorf("at least one filtering parameter must exist")
	}

	data, reqErr := h.do(GET, ClanEndpoint+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &ClanList); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &ClanList, nil
}

func (h *Client) GetClan(ClanTag string) (*Clan, error) {
	var Clan Clan
	ClanTag = string(toClanTag(ClanTag))

	data, reqErr := h.do(GET, ClanEndpoint+"/"+url.PathEscape(ClanTag), "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &Clan); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &Clan, nil
}

func (h *Client) GetClanMembers(ClanTag string) ([]ClanMember, error) {
	ClanMems := make([]ClanMember, 0)
	ClanTag = string(toClanTag(ClanTag))

	data, reqErr := h.do(GET, ClanEndpoint+"/"+url.PathEscape(ClanTag)+"/members/", "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &ClanMems); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return ClanMems, nil
}

func (h *Client) GetClanWarLog(ClanTag string) (*WarLog, error) {
	var ClanWarLog WarLog
	ClanTag = string(toClanTag(ClanTag))

	data, reqErr := h.do(GET, ClanEndpoint+"/"+url.PathEscape(ClanTag)+"/warlog/", "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &ClanWarLog); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &ClanWarLog, nil
}

func (h *Client) GetClanCurrentWar(ClanTag string) (*CurrentWar, error) {
	var ClanWar CurrentWar
	ClanTag = string(toClanTag(ClanTag))

	data, reqErr := h.do(GET, ClanEndpoint+"/"+url.PathEscape(ClanTag)+"/currentwar/", "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &ClanWar); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &ClanWar, nil
}

func (h *Client) GetClanWarLeagueGroup(ClanTag string) { //waiting for next cwl

}

func (h *Client) GetCWLWars(WarTag string) { //above

}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Player Methods
//_______________________________________________________________________

func (h *Client) GetPlayer(PlayerTag string) (*Player, error) {
	var Player Player
	PlayerTag = string(toPlayerTag(PlayerTag))

	data, reqErr := h.do(GET, PlayerEndpoint+"/"+url.PathEscape(PlayerTag), "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &Player); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &Player, nil
}

// Note: This is a custom method I made to make use of concurrency.
// A slice of players will be returned with no error, however players that had an error (i.e are banned or the tag never existed)
// will be nil inside the slice. The order of the player tags is kept intact.
func (h *Client) GetPlayers(PlayerTags []string) []*Player {
	Players := make([]*Player, len(PlayerTags))
	PlayerMap := make(map[string]*Player)

	for idx, _tag := range PlayerTags { // correct all tags to check with the map later
		PlayerTags[idx] = string(toPlayerTag(_tag))
	}

	var playerWg sync.WaitGroup
	var playerMapMutex sync.Mutex

	playerWg.Add(len(PlayerTags))
	for _, tag := range PlayerTags {
		go func(t string) {
			defer playerWg.Done()
			player, err := h.GetPlayer(t)
			if err != nil {
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
func (h *Client) VerifyPlayerToken(PlayerTag string, Token string) (*PlayerVerification, error) {
	var Verification PlayerVerification
	PlayerTag = string(toPlayerTag(PlayerTag))

	data, reqErr := h.do(POST, PlayerEndpoint+"/"+url.PathEscape(PlayerTag)+"/verifytoken/", fmt.Sprintf(`{"token": "%s"}`, Token), false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &Verification); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &Verification, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// League Methods
//_______________________________________________________________________

func (h *Client) GetLeagues(options *searchOptions) (*LeagueData, error) {
	var LeagueData LeagueData
	var opts string

	if options != nil {
		opts = options.ToQuery()
	}
	data, reqErr := h.do(GET, LeagueEndpoint+"/"+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &LeagueData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &LeagueData, nil
}

func (h *Client) GetLeague(LeagueID LeagueID) (*League, error) {
	var League League
	data, reqErr := h.do(GET, LeagueEndpoint+"/"+fmt.Sprint(LeagueID), "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &League); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &League, nil
}

// Ensure you pass in 29000022 for LeagueID
func (h *Client) GetLeagueSeasons(LeagueID LeagueID, options *searchOptions) (*SeasonData, error) {
	var SeasonData SeasonData
	var opts string
	if LeagueID != LegendLeague {
		fmt.Println(ErrInvalidLeague.Error())
		return nil, ErrInvalidLeague
	}
	if options != nil {
		opts = options.ToQuery()
	}

	data, reqErr := h.do(GET, LeagueEndpoint+"/"+fmt.Sprint(LeagueID)+"/seasons"+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &SeasonData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &SeasonData, nil
}

// Be cautious when using this. the data returned is massive. recommended to add args of {"limit": limit} and use the cursors for more data
func (h *Client) GetLeagueSeasonInfo(LeagueID LeagueID, SeasonID string, options *searchOptions) (*SeasonInfo, error) {
	var SeasonInfo SeasonInfo
	var opts string
	if LeagueID != LegendLeague {
		fmt.Println(ErrInvalidLeague)
		LeagueID = LegendLeague
	}

	match, matchErr := regexp.MatchString("^20[0-2][0-9]-((0[1-9])|(1[0-2]))$", SeasonID)
	if matchErr != nil {
		return nil, matchErr
	}

	if !match {
		return nil, fmt.Errorf("invalid season format, format must match the YYYY-MM date format")
	}

	if options != nil {
		opts = options.ToQuery()
	}

	data, reqErr := h.do(GET, LeagueEndpoint+"/"+fmt.Sprint(LeagueID)+"/seasons/"+SeasonID+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &SeasonInfo); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &SeasonInfo, nil
}

func (h *Client) GetWarLeagues(options *searchOptions) (*LeagueData, error) {
	var WarLeagueData LeagueData
	var opts string

	if options != nil {
		opts = options.ToQuery()
	}
	data, reqErr := h.do(GET, WarLeagueEndpoint+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &WarLeagueData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &WarLeagueData, nil
}

func (h *Client) GetWarLeague(WarLeagueID WarLeagueID) (*League, error) {
	var WarLeague League
	data, reqErr := h.do(GET, WarLeagueEndpoint+"/"+fmt.Sprint(WarLeagueID), "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &WarLeague); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &WarLeague, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Location Methods
//_______________________________________________________________________

//This should be passed ideally with nothing, kwargs aren't necessary here but only for the sake of completeness.
func (h *Client) GetLocations(options *searchOptions) (*LocationData, error) {
	var LocationData LocationData
	var opts string
	if options != nil {
		opts = options.ToQuery()
	}
	data, reqErr := h.do(GET, LocationEndpoint+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &LocationData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &LocationData, nil
}

func (h *Client) GetLocation(LocationID LocationID) (*Location, error) {
	var Location Location
	data, reqErr := h.do(GET, LocationEndpoint+"/"+fmt.Sprint(LocationID), "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &Location); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &Location, nil
}

// Main Village Clan Rankings
func (h *Client) GetLocationClans(LocationID LocationID, options *searchOptions) (*ClanRankingList, error) {
	var ClanData ClanRankingList
	var opts string
	if options != nil {
		opts = options.ToQuery()
	}
	data, reqErr := h.do(GET, LocationEndpoint+"/"+fmt.Sprint(LocationID)+"/rankings/clans"+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &ClanData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &ClanData, nil
}

// Builder Hall Clan Rankings
func (h *Client) GetLocationClansVersus(LocationID LocationID, options *searchOptions) (*ClanVersusRankingList, error) {
	var ClanVersusData ClanVersusRankingList
	var opts string
	if options != nil {
		opts = options.ToQuery()
	}
	data, reqErr := h.do(GET, LocationEndpoint+"/"+fmt.Sprint(LocationID)+"/rankings/clans-versus"+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &ClanVersusData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &ClanVersusData, nil
}

// Main Village Player Rankings
func (h *Client) GetLocationPlayers(LocationID LocationID, options *searchOptions) (*PlayerRankingList, error) {
	var PlayerData PlayerRankingList
	var opts string
	if options != nil {
		opts = options.ToQuery()
	}
	data, reqErr := h.do(GET, LocationEndpoint+"/"+fmt.Sprint(LocationID)+"/rankings/players"+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &PlayerData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &PlayerData, nil
}

// Builder Hall Player Rankings
func (h *Client) GetLocationPlayersVersus(LocationID LocationID, options *searchOptions) (*PlayerVersusRankingList, error) {
	var PlayerVersusData PlayerVersusRankingList
	var opts string
	if options != nil {
		opts = options.ToQuery()
	}
	data, reqErr := h.do(GET, LocationEndpoint+"/"+fmt.Sprint(LocationID)+"/rankings/players-versus"+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &PlayerVersusData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &PlayerVersusData, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Gold Pass Method
//_______________________________________________________________________

func (h *Client) GetGoldPass() (*GoldPassSeason, error) {
	var GoldPass GoldPassSeason
	data, reqErr := h.do(GET, GoldpassEndpoint, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &GoldPass); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &GoldPass, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Label Methods
//_______________________________________________________________________

func (h *Client) GetClanLabels(options *searchOptions) (*LabelsData, error) {
	var LabelsData LabelsData
	var opts string
	if options != nil {
		opts = options.ToQuery()
	}
	data, reqErr := h.do(GET, LabelEndpoint+"/clans"+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &LabelsData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &LabelsData, nil
}

func (h *Client) GetPlayerLabels(options *searchOptions) (*LabelsData, error) {
	var LabelsData LabelsData
	var opts string
	if options != nil {
		opts = options.ToQuery()
	}
	data, reqErr := h.do(GET, LabelEndpoint+"/players"+opts, "", false)
	if reqErr != nil {
		return nil, reqErr
	}

	var err APIError
	if jsonErr := json.Unmarshal(data, &LabelsData); jsonErr != nil {
		jsonErr2 := json.Unmarshal(data, &err)
		if jsonErr2 != nil {
			return nil, jsonErr2
		}
		return nil, &err
	}
	return &LabelsData, nil
}
