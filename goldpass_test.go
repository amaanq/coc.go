package coc

import "testing"

func Test_GoldPass(t *testing.T) {
	init_dummy()

	goldPass, err := dummyClient.GetGoldPass()
	if err != nil {
		t.Errorf("GetGoldPass() error = %v", err)
	}
	t.Log("gold pass", goldPass)
}
