package location

type PlayerData struct {
	Players []Player `json:"items,omitempty"`
	Paging  struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type PlayerVersusData struct {
	PlayersVersus []PlayerVersus `json:"items,omitempty"`
	Paging        struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type Player struct {
	Tag          string     `json:"tag,omitempty"`
	Name         string     `json:"name,omitempty"`
	ExpLevel     int64      `json:"expLevel,omitempty"`
	Trophies     int64      `json:"trophies,omitempty"`
	AttackWins   int64      `json:"attackWins,omitempty"`
	DefenseWins  int64      `json:"defenseWins,omitempty"`
	Rank         int64      `json:"rank,omitempty"`
	PreviousRank int64      `json:"previousRank,omitempty"`
	Clan         PlayerClan `json:"clan,omitempty"`
	League       League     `json:"league,omitempty"`
}

type PlayerVersus struct {
	Tag              string     `json:"tag,omitempty"`
	Name             string     `json:"name,omitempty"`
	ExpLevel         int64      `json:"expLevel,omitempty"`
	VersusTrophies   int64      `json:"versusTrophies,omitempty"`
	VersusBattleWins int64      `json:"versusBattleWins,omitempty"`
	Rank             int64      `json:"rank,omitempty"`
	PreviousRank     int64      `json:"previousRank,omitempty"`
	Clan             PlayerClan `json:"clan,omitempty"`
	League           League     `json:"league,omitempty"`
}

type PlayerClan struct {
	Tag       string    `json:"tag,omitempty"`
	Name      string    `json:"name,omitempty"`
	BadgeUrls BadgeUrls `json:"badgeUrls,omitempty"`
}

type League struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	IconUrls struct {
		Small  string `json:"small,omitempty"`
		Tiny   string `json:"tiny,omitempty"`
		Medium string `json:"medium,omitempty"`
	} `json:"iconUrls,omitempty"`
}
