package coc

import "testing"

func Test_GetPlayer(t *testing.T) {
	init_dummy()

	player, err := dummyClient.GetPlayer("#2PP")
	if err != nil {
		t.Errorf("GetPlayer() error = %v", err)
	}
	t.Log("player", player)
}

func Test_GetPlayers(t *testing.T) {
	init_dummy()

	players := dummyClient.GetPlayers([]string{"#2PP", "#8GG", "#9UU", "#Y98", "#LQL"})
	t.Log("players", players)
}
