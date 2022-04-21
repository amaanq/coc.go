package coc

import "testing"

func Test_GetPlayerLabels(t *testing.T) {
	init_dummy()

	labels, err := dummyClient.GetPlayerLabels(nil)
	if err != nil {
		t.Errorf("GetPlayerLabels() error = %v", err)
	}
	t.Log("player labels", labels)
}

func Test_GetClanLabels(t *testing.T) {
	init_dummy()

	labels, err := dummyClient.GetClanLabels(nil)
	if err != nil {
		t.Errorf("GetClanLabels() error = %v", err)
	}
	t.Log("clan labels", labels)
}
