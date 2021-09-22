package main

import (
	//"fmt"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/amaanq/coc.go/client"
)


func main() {
	t := time.Now()
	godotenv.Load(".env")
	H := client.Initialize(os.Getenv("email"), os.Getenv("password"))

	err := H.APILopin()
	if err != nil {
		panic(err)
	}
	err = H.GetKeys()
	if err != nil {
		panic(err)
	}
	err = H.AddOrDeleteKeysAsNecessary()
	if err != nil {
		panic(err)
	}
	cln, err := H.GetLocationPlayersVersus("32000249")
	if err != nil {
		panic(err)
	}
	for _, mem := range cln.PlayersVersus {
		fmt.Println("data:", mem.Name, mem.ExpLevel, mem.Tag, mem.VersusTrophies)
	}
	t2 := time.Since(t)
	fmt.Printf("Took %f seconds\n", t2.Seconds())
}
