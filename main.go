package main

import (
	"fmt"
	"os"

	"github.com/amaanq/coc.go/client"
	"github.com/joho/godotenv"
)

func New(credentials map[string]string) *client.HTTPSessionManager {
	return client.Initialize(credentials)
}

func main() {
	godotenv.Load(".env")

	client := New(map[string]string{os.Getenv("email"): os.Getenv("password"), os.Getenv("email2"): os.Getenv("password2"), os.Getenv("email3"): os.Getenv("password3")})

	player, err := client.GetPlayer("#2PP")
	if err.Err() != nil {
		fmt.Println(err.Err().Error())
		panic(err.Err())
	}
	fmt.Println(player.Name)

	clan, err := client.GetClan("#2PP")
	if err.Err() != nil {
		panic(err)
	}
	fmt.Println(clan.Name)

	list, clientErr := client.SearchClans(map[string]string{"minClanLevel": "20"})
	if clientErr.Err() != nil {
		fmt.Println("msg", clientErr.Message)
		panic(clientErr.Err())
	}

	fmt.Println(list.Clans[0])
}