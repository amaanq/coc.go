package coc

import (
	"strings"
	"testing"
)

func Test_GetClanSearch(t *testing.T) {
	init_dummy()
	clans, err := dummyClient.SearchClans(ClanSearchOptions().SetName("test").SetLimit(10).SetMaxMembers(40))
	if err != nil {
		t.Errorf("GetClan() error = %v", err)
	}
	if len(clans.Clans) != 10 {
		t.Errorf("GetClan() = %v, want %v", len(clans.Clans), 10)
	}
	for _, clan := range clans.Clans {
		if strings.ToLower(clan.Name) != "test" {
			t.Errorf("GetClan() = %v, want %v", clan.Name, "test")
		}
	}
}
