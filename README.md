[![Build Status](https://app.travis-ci.com/amaanq/coc.go.svg?branch=master)](https://app.travis-ci.com/amaanq/coc.go.svg?branch=master)
[![codecov](https://codecov.io/gh/amaanq/coc.go/branch/master/graph/badge.svg)](https://codecov.io/gh/amaanq/coc.go)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/amaanq/coc.go)
[![Go Report Card](https://goreportcard.com/badge/github.com/amaanq/coc.go)](https://goreportcard.com/report/github.com/amaanq/coc.go)


## Getting Started

### Installing

#### This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` *will always pull the latest tagged release from the master branch.*

```sh
go get github.com/amaanq/coc.go
```
</br>

## Usage

Import the package into your project.

```go
import "github.com/amaanq/coc.go"
```

Construct a new Clash API Client which can be used to access the variety of 
Clash API functions. Please ENSURE your credentials are valid and please DO NOT use a password you use for important credentials, 
even though nothing is logged or stored here. Initialize automatically logs into your developer account, checks your keys, and adds or deletes them as necessary
depending on your IP. You are able to login with multiple credentials in the event your application/project is large and draws a lot of users, key handling is done automatically for all logins. 

```go
package main

import (
    "github.com/amaanq/coc.go"
)

func main() {
    APIClient, err := coc.New(map[string]string{"email": "password", "email2": "password2", "email3": "password3"})
    if err != nil {
        // handle err 
    }
}
```
You can add as many logins as you want, but it'll be slower with more logins. 
I recommend no more than 3.
</br></br>

## See Documentation and /examples below for more detailed information.

## **NOTICE**: This library and the Clash API are unfinished.
Because of that there may be major changes to library in the future.

The coc.go code is fairly well documented at this point and is currently
the only documentation available. 
There are 4 main types of endpoints for the API. Player, Clan, Location, and League. Minor ones are label and goldpass.
At the moment the CWL endpoints have yet to be implemented since I don't have sample json to base the structs off of yet. This will be done next cwl. 


Here's how you can fetch player data and display it to your terminal.
```go
import (
    "github.com/amaanq/coc.go"
)
// Obviously load a client first...
func main() {
  player, err := APIClient.GetPlayer("#YourTag")
  if err != nil {
    // handle err 
  }

  fmt.Printf("Player: %+v\n", player)
  fmt.Println("My name is: ", player.Name)

  //Same for a clan: 

  clan, err := APIClient.GetClan("#YourTag")
  if err != nil {
    // handle err 
  }

  fmt.Printf("Clan: %+v\n", clan)
  fmt.Println("My clan name is", clan.Name,"and we have", clan.Members, "members in our clan. We have won", clan.WarWins, "wars so come join us!\nThese are our members:")
  for idx, member := range clan.MemberList {
    fmt.Printf("[%d]: %s (%s)\n", idx, member.Name, member.Role)
  }
}
```
</br>

## **IMPORTANT**: There are some endpoints that can pass in a variety of arguments (i.e searching for clans). 
### This variety of arguments generally allows a limit, before, or after parameter to be set. However for clans there are a lot more options.
### So, you can make use SearchOptions or ClanSearchOptions accordingly. See example below:
</br>

## ANOTHER NOTE: You cannot set `before` and `after` at the same time. The first one set will be used.

```go
import (
    "github.com/amaanq/coc.go"
)

// Obviously load a client first...
func main() {
  clans, err := APIClient.SearchClans(coc.ClanSearchOptions().SetName("test").SetLimit(10).SetMaxMembers(40))
  // which is the same as this (different order of setting)
  // clans, err := APIClient.SearchClans(coc.ClanSearchOptions().SetName("test").SetMaxMembers(40).SetLimit(10))
  if err != nil {
      // handle err 
  }

  for _, clan := range clans.Clans {
      fmt.Println("data:", clan.Name, clan.RequiredTownhallLevel, clan.ClanLevel, clan.RequiredTrophies)
      if clan.ClanLevel >= 15 {
          fmt.Println("Found a clan you're looking for!", clan.Name, clan.Tag)
          break
      }
  }

  leagues, err := APIClient.GetLeagues(coc.SearchOptions().SetLimit(10).SetAfter(2))
  if err != nil {
    // handle err 
  }
}
```
You can enter the map[string]string args as one map with several key-value pairs, or several maps with 1 key-value pair; this parsing is handled automatically.
Enter every data type as a string, even ints. You can also leave the map empty, this is acceptable but some functions will error as one argument is required (SearchClans being one).


## Subdirectories

* [examples](./examples)

---
