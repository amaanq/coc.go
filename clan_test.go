package coc

import "testing"

func Test_GetClan(t *testing.T) {
	// make a test for clan 2pp
	init_dummy()
	clan, err := dummyClient.GetClan("2pp")
	if err != nil {
		t.Errorf("GetClan() error = %v", err)
	}
	if clan.Tag != "#2PP" {
		t.Errorf("GetClan() = %v, want %v", clan.Tag, "#2PP")
	}
	if clan.Name != "The Order" {
		t.Errorf("GetClan() = %v, want %v", clan.Name, "The Order")
	}
	t.Log(clan)
}
