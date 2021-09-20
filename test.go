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
	fmt.Println("logged in")
	err = H.GetKeys()
	if err != nil {
		panic(err)
	}
	fmt.Println("got keys")
	err = H.AddOrDeleteKeysAsNecessary()
	if err != nil {
		panic(err)
	}
	fmt.Println("finished adding keys")
	t2 := time.Since(t)
	fmt.Printf("Took %f seconds\n", t2.Seconds())
	H.ViewData()
}
