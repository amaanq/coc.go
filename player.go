package coc

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Player
//_______________________________________________________________________

type Player struct {
	League               League           `json:"league"`
	Role                 string           `json:"role"`
	Name                 string           `json:"name"`
	Tag                  string           `json:"tag"`
	WarPreference        string           `json:"warPreference"`
	Spells               []Hero           `json:"spells"`
	Heroes               []Hero           `json:"heroes"`
	Troops               []Hero           `json:"troops"`
	Labels               []Label          `json:"labels"`
	Achievements         []Achievement    `json:"achievements"`
	LegendStatistics     LegendStatistics `json:"legendStatistics"`
	Clan                 Clan             `json:"clan"`
	BestTrophies         int              `json:"bestTrophies"`
	DefenseWins          int              `json:"defenseWins"`
	BestVersusTrophies   int              `json:"bestVersusTrophies"`
	VersusTrophies       int              `json:"versusTrophies"`
	Donations            int              `json:"donations"`
	DonationsReceived    int              `json:"donationsReceived"`
	BuilderHallLevel     int              `json:"builderHallLevel"`
	VersusBattleWINS     int              `json:"versusBattleWins"`
	AttackWins           int              `json:"attackWins"`
	WarStars             int              `json:"warStars"`
	VersusBattleWinCount int              `json:"versusBattleWinCount"`
	Trophies             int              `json:"trophies"`
	ExpLevel             int              `json:"expLevel"`
	TownHallWeaponLevel  int              `json:"townHallWeaponLevel"`
	TownHallLevel        int              `json:"townHallLevel"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Hero
//_______________________________________________________________________

type Hero struct {
	Name               string  `json:"name"`
	Village            Village `json:"village"`
	Level              int     `json:"level"`
	MaxLevel           int     `json:"maxLevel"`
	SuperTroopIsActive bool    `json:"superTroopIsActive,omitempty"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Legend League
//_______________________________________________________________________

type LegendStatistics struct {
	PreviousSeason   Season        `json:"previousSeason"`
	BestSeason       Season        `json:"bestSeason"`
	BestVersusSeason Season        `json:"bestVersusSeason"`
	CurrentSeason    CurrentSeason `json:"currentSeason"`
	LegendTrophies   int           `json:"legendTrophies"`
}

type Season struct {
	ID       string `json:"id,omitempty"`
	Rank     int    `json:"rank,omitempty"`
	Trophies int    `json:"trophies,omitempty"`
}

type CurrentSeason struct {
	Rank     int `json:"rank,omitempty"`
	Trophies int `json:"trophies,omitempty"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Village
//_______________________________________________________________________

type Village string

const (
	BuilderBase Village = "builderBase"
	Home        Village = "home"
)

type PlayerVerification struct {
	Tag    string `json:"tag"`
	Token  string `json:"token"`
	Status string `json:"status"`
}
