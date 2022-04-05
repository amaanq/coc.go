package main

import (
	"fmt"

	"github.com/amaanq/coc.go"
)

func main() {
	client, cErr := coc.New(map[string]string{"username": "password"})
	if cErr != nil {
		panic(cErr)
	}

	player, err := client.GetPlayer("#2PP")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Player name: ", player.Name)
}
