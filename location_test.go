package coc

import "testing"

func Test_GetLocations(t *testing.T) {
	init_dummy()

	locations, err := dummyClient.GetLocations(nil)
	if err != nil {
		t.Errorf("GetLocations() error = %v", err)
	}
	t.Log("locations", locations)
}

func Test_GetLocationID(t *testing.T) {
	init_dummy()

	location, err := dummyClient.GetLocation(UnitedStates)
	if err != nil {
		t.Errorf("GetLocation() error = %v", err)
	}
	t.Log("location", location)
}
