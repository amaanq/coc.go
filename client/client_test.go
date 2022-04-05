package client

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var ( // Bad practice but oh well I'm not loading the client for every single test..
	_              = godotenv.Load("../.env")
	DummyClient, _ = New(map[string]string{os.Getenv("email"): os.Getenv("password")})
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func TestHTTPSessionManager_GetPlayer(t *testing.T) {
	type args struct {
		PlayerTag string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "First", args: args{PlayerTag: "#2PP"}, wantErr: false},
		{name: "Second", args: args{PlayerTag: "#2PP"}, wantErr: false},
		{name: "Third", args: args{PlayerTag: "#2PP"}, wantErr: false}, // Timings should be instant..
		{name: "New", args: args{PlayerTag: "#8GG"}, wantErr: false},   // Back to slow
		{name: "Bad Tag", args: args{PlayerTag: "#222"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer duration(track("GET Player " + tt.args.PlayerTag + " Time"))
			_, err := DummyClient.GetPlayer(tt.args.PlayerTag)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPSessionManager.GetPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHTTPSessionManager_SearchClans(t *testing.T) {
	type args struct {
		options *clanSearchOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test", args: args{options: ClanSearchOptions().SetName("test").SetLimit(10).SetMaxMembers(40)}, wantErr: false},
		{name: "Test Caching", args: args{options: ClanSearchOptions().SetName("test").SetLimit(10).SetMaxMembers(40)}, wantErr: false},
		{name: "Test Caching Unordered", args: args{options: ClanSearchOptions().SetName("test").SetMaxMembers(40).SetLimit(10)}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer duration(track("GET Search Clans " + tt.args.options.ToQuery() + " Time"))
			_, err := DummyClient.SearchClans(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPSessionManager.SearchClans() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestHTTPSessionManager_GetLeagues(t *testing.T) {
	type args struct {
		options *searchOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test", args: args{options: SearchOptions().SetLimit(10).SetAfter(2)}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer duration(track("GET Leagues " + tt.args.options.ToQuery() + " Time"))
			_, err := DummyClient.GetLeagues(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPSessionManager.GetLeagues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
