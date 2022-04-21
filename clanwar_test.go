package coc

import (
	"testing"
)

func Test_GetClanCurrentWar(t *testing.T) {
	init_dummy()
	clanWar, err := dummyClient.GetClanCurrentWar("LCG8C2CQ")
	if err != nil {
		t.Errorf("GetClanCurrentWar() error = %v", err)
	}

	t.Log(clanWar.PreparationStartTime())
	t.Log(clanWar.StartTime())
	t.Log(clanWar.EndTime())
	t.Log(clanWar.State)
	t.Log(clanWar.TeamSize)
}
