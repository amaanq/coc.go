package main

import (
	"encoding/json"
	"fmt"
	//"math/rand"
	"os"
	//"time"

	"github.com/amaanq/coc.go/client"
	"github.com/joho/godotenv"
)

func main() {
	//t := time.Now()
	godotenv.Load(".env")
	H := client.Initialize(map[string]string{os.Getenv("email"): os.Getenv("password"), os.Getenv("email2"): os.Getenv("password2"), os.Getenv("email3"): os.Getenv("password3")})

	p, _ := H.GetClanCurrentWar("#2LU99C90U")
	
	for _, n := range p.Clan.Members {
		fmt.Println(n)
	}
	fmt.Println(p)
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}
