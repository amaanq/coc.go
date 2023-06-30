package coc

import "testing"

func Test_GetLeagues(t *testing.T) {
	init_dummy()

	leagues, err := dummyClient.GetLeagues(nil)
	if err != nil {
		t.Errorf("GetLeagues() error = %v", err)
	}
	t.Log("leagues", leagues)
}

func Test_GetLeagueID(t *testing.T) {
	init_dummy()

	league, err := dummyClient.GetLeague(ChampionLeagueI)
	if err != nil {
		t.Errorf("GetLeague() error = %v", err)
	}
	t.Log("league", league)
}

func Test_GetWarLeagues(t *testing.T) {
	init_dummy()

	warLeagues, err := dummyClient.GetWarLeagues(nil)
	if err != nil {
		t.Errorf("GetWarLeagues() error = %v", err)
	}
	t.Log("war leagues", warLeagues)
}

func Test_GetWarLeagueID(t *testing.T) {
	init_dummy()

	warLeague, err := dummyClient.GetWarLeague(ChampionWarLeagueI)
	if err != nil {
		t.Errorf("GetWarLeague() error = %v", err)
	}
	t.Log("war league", warLeague)
}

func Test_GetLeagueSeasons(t *testing.T) {
	init_dummy()

	seasons, err := dummyClient.GetLeagueSeasons(LegendLeague, nil)
	if err != nil {
		t.Errorf("GetLeagueSeasons() error = %v", err)
	}
	t.Log("seasons", seasons)

	seasons, err = dummyClient.GetLeagueSeasons(CrystalLeagueI, nil)
	if err == nil {
		t.Errorf("GetLeagueSeasons() error = %v", seasons)
	}
	t.Log("valid err", err)
}

func Test_GetLeagueSeasonID(t *testing.T) {
	init_dummy()

	season, err := dummyClient.GetLeagueSeasonInfo(LegendLeague, "2020-12", nil)
	if err != nil {
		t.Errorf("GetLeagueSeasonInfo() error = %v", err)
	}
	t.Log("season", season)
}
