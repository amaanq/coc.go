package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/amaanq/coc.go/client"
	"github.com/joho/godotenv"
)

func main() {
	t := time.Now()
	godotenv.Load(".env")
	H := client.Initialize(map[string]string{os.Getenv("email"): os.Getenv("password"), os.Getenv("email2"): os.Getenv("password2"), os.Getenv("email3"): os.Getenv("password3")})

	for i := 0; i < 100000; i++ {
		if i%100 == 0 {
			mp := map[string]string{"name": "the" + string(rune(rand.Intn(90)+32))}
			fmt.Println(mp)
			cln, err := H.SearchClans(mp)
			if err != nil {
				panic(err)
			}
			for _, mem := range cln.Clans {
				fmt.Println("data:", mem.Name, mem.RequiredTownhallLevel, mem.ClanLevel, mem.RequiredTrophies)
			}
			t2 := time.Since(t)
			fmt.Printf("NEW: Took %f seconds\n", t2.Seconds())
			t = time.Now()
		}
		cln, err := H.SearchClans(map[string]string{"name": "hey"})
		if err != nil {
			panic(err)
		}
		for range cln.Clans {
			//fmt.Println("data:", mem.Name, mem.RequiredTownhallLevel, mem.ClanLevel, mem.RequiredTrophies)
		}
		t2 := time.Since(t)
		fmt.Printf("Took %f seconds\n", t2.Seconds())
		t = time.Now()
	}
}
