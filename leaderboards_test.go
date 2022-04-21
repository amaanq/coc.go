package coc

import "testing"

func Test_GetClanRanking(t *testing.T) {
	init_dummy()

	ranking, err := dummyClient.GetLocationClans(UnitedStates, nil)
	if err != nil {
		t.Errorf("GetClanRanking() error = %v", err)
	}
	t.Log("clan ranking", ranking.Clans)
}

func Test_GetClanVersusRanking(t *testing.T) {
	init_dummy()

	ranking, err := dummyClient.GetLocationClansVersus(UnitedStates, nil)
	if err != nil {
		t.Errorf("GetClanVersusRanking() error = %v", err)
	}
	t.Log("clan versus ranking", ranking.ClansVersus)
}

func Test_GetPlayerRanking(t *testing.T) {
	init_dummy()

	ranking, err := dummyClient.GetLocationPlayers(UnitedStates, nil)
	if err != nil {
		t.Errorf("GetPlayerRanking() error = %v", err)
	}
	t.Log("player ranking", ranking.Players)
}

func Test_GetPlayerVersusRanking(t *testing.T) {
	init_dummy()

	ranking, err := dummyClient.GetLocationPlayersVersus(UnitedStates, nil)
	if err != nil {
		t.Errorf("GetPlayerVersusRanking() error = %v", err)
	}
	t.Log("player versus ranking", ranking.PlayersVersus)
}
