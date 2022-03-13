package client

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestClient(t *testing.T) {
	godotenv.Load("../.env")

	H, Herr := New(map[string]string{os.Getenv("email"): os.Getenv("password"), os.Getenv("email2"): os.Getenv("password2"), os.Getenv("email3"): os.Getenv("password3")})
	if Herr != nil {
		panic(Herr)
	}

	defer duration(track("Login Time"))

	fmt.Println(H.logins)

	player, err := H.GetPlayer("#2PP")
	if err.Err() != nil {
		fmt.Println(err.Err().Error())
		panic(err.Err())
	}
	fmt.Println(player.Name)

	clan, err := H.GetClan("#2PP")
	if err.Err() != nil {
		panic(err)
	}
	fmt.Println(clan.Name)

	list, clientErr := H.SearchClans(map[string]string{"minClanLevel": "20"})
	if clientErr.Err() != nil {
		fmt.Println("msg", clientErr.Message)
		panic(clientErr.Err())
	}

	fmt.Println(list.Clans[0])

	fmt.Println(len(H.allKeys.Keys))

	players := H.GetPlayers([]string{"#2PP", "#8GG"})
	fmt.Println(players[0].Achievements, players[1].BestTrophies)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
