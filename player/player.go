package player

import "github.com/amaanq/coc.go/tag"

type Player struct {
	Tag                  tag.PlayerTag    `json:"tag"`
	Name                 string           `json:"name"`
	TownHallLevel        int              `json:"townHallLevel"`
	TownHallWeaponLevel  int              `json:"townHallWeaponLevel"`
	ExpLevel             int              `json:"expLevel"`
	Trophies             int              `json:"trophies"`
	BestTrophies         int              `json:"bestTrophies"`
	WarStars             int              `json:"warStars"`
	AttackWins           int              `json:"attackWins"`
	DefenseWins          int              `json:"defenseWins"`
	BuilderHallLevel     int              `json:"builderHallLevel"`
	VersusTrophies       int              `json:"versusTrophies"`
	BestVersusTrophies   int              `json:"bestVersusTrophies"`
	VersusBattleWINS     int              `json:"versusBattleWins"`
	Role                 string           `json:"role"`
	WarPreference        string           `json:"warPreference"`
	Donations            int              `json:"donations"`
	DonationsReceived    int              `json:"donationsReceived"`
	Clan                 Clan             `json:"clan"`
	League               League           `json:"league"`
	LegendStatistics     LegendStatistics `json:"legendStatistics"`
	Achievements         []Achievement    `json:"achievements"`
	VersusBattleWinCount int              `json:"versusBattleWinCount"`
	Labels               []Label          `json:"labels"`
	Troops               []Hero           `json:"troops"`
	Heroes               []Hero           `json:"heroes"`
	Spells               []Hero           `json:"spells"`
}

type Achievement struct {
	Name           string  `json:"name"`
	Stars          int     `json:"stars"`
	Value          int     `json:"value"`
	Target         int     `json:"target"`
	Info           string  `json:"info"`
	CompletionInfo string  `json:"completionInfo"`
	Village        Village `json:"village"`
}

type Clan struct {
	Tag       tag.ClanTag `json:"tag"`
	Name      string      `json:"name"`
	ClanLevel int         `json:"clanLevel"`
	BadgeUrls BadgeUrls   `json:"badgeUrls"`
}

type BadgeUrls struct {
	Small  string `json:"small"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type Hero struct {
	Name               string  `json:"name"`
	Level              int     `json:"level"`
	MaxLevel           int     `json:"maxLevel"`
	Village            Village `json:"village"`
	SuperTroopIsActive bool    `json:"superTroopIsActive,omitempty"`
}

type Label struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	IconUrls LabelIconUrls `json:"iconUrls"`
}

type LabelIconUrls struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
}

type League struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	IconUrls LeagueIconUrls `json:"iconUrls"`
}

type LeagueIconUrls struct {
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
	Medium string `json:"medium"`
}

type LegendStatistics struct {
	LegendTrophies   int           `json:"legendTrophies"`
	PreviousSeason   Season        `json:"previousSeason"`
	BestSeason       Season        `json:"bestSeason"`
	BestVersusSeason Season        `json:"bestVersusSeason"`
	CurrentSeason    CurrentSeason `json:"currentSeason"`
}

type Season struct {
	ID       string `json:"id"`
	Rank     int    `json:"rank"`
	Trophies int    `json:"trophies"`
}

type CurrentSeason struct {
	Rank     int `json:"rank"`
	Trophies int `json:"trophies"`
}

type Village string

const (
	BuilderBase Village = "builderBase"
	Home        Village = "home"
)

type Verification struct {
	Tag    string `json:"tag"`
	Token  string `json:"token"`
	Status string `json:"status"`
}
