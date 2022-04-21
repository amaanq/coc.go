package coc

import (
	"testing"
)

func TestGetAchievement(t *testing.T) {
	init_dummy()
	type args struct {
		tag               string
		wantedAchievement Achievement
	}
	tests := []struct {
		name    string
		args    args
		want    Achievement
		wantErr bool
	}{
		{name: "2PP", args: args{tag: "#2PP", wantedAchievement: KeepYourAccountSafeOld}, wantErr: false},
		{name: "2PP", args: args{tag: "#2PP", wantedAchievement: KeepYourAccountSafeSCID}, wantErr: false},
		{name: "2PP", args: args{tag: "#2PP", wantedAchievement: ShatteredAndScattered}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player, err := dummyClient.GetPlayer(tt.args.tag)
			if err != nil {
				t.Errorf("GetPlayer() error = %v", err)
				return
			}
			got, err := GetAchievement(player.Achievements, tt.args.wantedAchievement)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAchievement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%s's matched Achievement: %v", player.Name, got)
		})
	}
}
