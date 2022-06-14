package coc

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Player
//_______________________________________________________________________

type Player struct {
	Tag                  PlayerTag        `json:"tag"`
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

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Hero
//_______________________________________________________________________

type Hero struct {
	Name               string  `json:"name"`
	Level              int     `json:"level"`
	MaxLevel           int     `json:"maxLevel"`
	Village            Village `json:"village"`
	SuperTroopIsActive bool    `json:"superTroopIsActive,omitempty"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Legend League
//_______________________________________________________________________

type LegendStatistics struct {
	LegendTrophies   int           `json:"legendTrophies"`
	PreviousSeason   Season        `json:"previousSeason"`
	BestSeason       Season        `json:"bestSeason"`
	BestVersusSeason Season        `json:"bestVersusSeason"`
	CurrentSeason    CurrentSeason `json:"currentSeason"`
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
	Tag    PlayerTag `json:"tag"`
	Token  string `json:"token"`
	Status string `json:"status"`
}
