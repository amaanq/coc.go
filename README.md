[![Build Status](https://app.travis-ci.com/amaanq/coc.go.svg?branch=master)](https://app.travis-ci.com/amaanq/coc.go.svg?branch=master)
[![codecov](https://codecov.io/gh/./branch/master/graph/badge.svg)](https://codecov.io/gh/github.com/amaanq/coc.go)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/amaanq/coc.go)
[![Go Report Card](https://goreportcard.com/badge/github.com/amaanq/coc.go)](https://goreportcard.com/report/github.com/amaanq/coc.go)


## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` *will always pull the latest tagged release from the master branch.*

```sh
go get github.com/amaanq/coc.go
```

### Usage

Import the package into your project.

```go
import "github.com/amaanq/coc.go"
```

Construct a new Clash API Client which can be used to access the variety of 
Clash API functions. Please ENSURE your credentials are valid and please DO NOT use a password you use for important credentials, 
even though nothing is logged or stored here. Initialize automatically logs into your developer account, checks your keys, and adds or deletes them as necessary
depending on your IP.

```go
ClashClient := client.Initialize("your email", "your password")
```

See Documentation and Examples below for more detailed information.

**NOTICE**: This library and the Clash API are unfinished.
Because of that there may be major changes to library in the future.

The coc.go code is fairly well documented at this point and is currently
the only documentation available. 
There are 4 main types of endpoints for the API. Player, Clan, Location, and League. Minor ones are label and goldpass.
At the moment the CWL endpoints have yet to be implemented since I don't have sample json to base the structs off of yet. This will be done next cwl. 

Here's how you can fetch player data and display it to your terminal.
```go
player, err := ClashClient.GetPlayer("#YourTag")
if err != nil {
  panic(err) // or fmt.Println(err.Error()) and return err
}
fmt.Printf("Player: %+v\n", player)
```

Same for a clan: 
```go
clan, err := ClashClient.GetClan("#YourTag")
if err != nil {
  panic(err) // or fmt.Println(err.Error()) and return err
}
fmt.Printf("Clan: %+v\n", player)
fmt.Println("My clan name is", clan.Name,"and we have", clan.Members, "members in our clan. We have won", clan.WarWins, "wars so come join us!\nThese are our members:")
for idx, member := range clan.MemberList {
  fmt.Printf("[%d]: %s (%s)", idx, member.Name, member.Role)
}
```

## Short-Hand Links to each Package in this Module

* [clan](./clan)

* [client](./client)

* [labels](./labels)

* [league](./league)

* [location](./location)

* [player](./player)

---
