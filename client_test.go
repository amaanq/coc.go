package coc

import (
	"log"
	"os"

	"testing"
	"time"

	"github.com/joho/godotenv"
)

var (
	dummyClient *Client
	dummyLoaded bool
)

// Load on every test if not loaded to ensure we have a test/dummy client to use for the tests
func init_dummy() {
	if dummyLoaded {
		return
	}
	dummyLoaded = true

	var err error
	godotenv.Load()
	dummyClient, err = New(map[string]string{os.Getenv("email"): os.Getenv("password")})
	if err != nil {
		log.Fatal(err)
	}
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func TestClient_GetPlayer(t *testing.T) {
	init_dummy()
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
			_, err := dummyClient.GetPlayer(tt.args.PlayerTag)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_SearchClans(t *testing.T) {
	init_dummy()
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
			_, err := dummyClient.SearchClans(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SearchClans() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestClient_GetLeagues(t *testing.T) {
	init_dummy()
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
			_, err := dummyClient.GetLeagues(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetLeagues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
