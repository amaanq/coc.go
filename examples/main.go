package main

import (
	"github.com/amaanq/coc.go"
)

func main() {
	client, err := coc.New(map[string]string{"username": "password"})
	if err != nil {
		panic(err)
	}

	player, err := client.GetPlayer("#2PP")
}
