package coc

import "time"

const (
	TimestampFormat = "20060102T150405.000Z"
)

type WarLog struct {
	Log []War `json:"items"`
}

type War struct {
	Result           Result  `json:"result"`
	Start            string  `json:"startTime"`
	End              string  `json:"endTime"`
	Clan             WarClan `json:"clan"`
	Opponent         WarClan `json:"opponent"`
	TeamSize         int     `json:"teamSize"`
	AttacksPerMember int     `json:"attacksPerMember"`
}

type WarClan struct {
	BadgeURLs             IconURLs    `json:"badgeUrls"`
	Tag                   string      `json:"tag"`
	Name                  string      `json:"name"`
	Members               []WarMember `json:"members"`
	ClanLevel             int         `json:"clanLevel"`
	Attacks               int         `json:"attacks"`
	Stars                 int         `json:"stars"`
	DestructionPercentage float64     `json:"destructionPercentage"`
	ExpEarned             int         `json:"expEarned"`
}

type CurrentWar struct {
	State            string  `json:"state"`
	PreparationStart string  `json:"preparationStartTime"`
	Start            string  `json:"startTime"`
	End              string  `json:"endTime"`
	Clan             WarClan `json:"clan"`
	Opponent         WarClan `json:"opponent"`
	TeamSize         int     `json:"teamSize"`
	AttacksPerMember int     `json:"attacksPerMember"`
}

type WarMember struct {
	Tag                string      `json:"tag"`
	Name               string      `json:"name"`
	Attacks            []WarAttack `json:"attacks"`
	BestOpponentAttack WarAttack   `json:"bestOpponentAttack"`
	TownhallLevel      int         `json:"townhallLevel"`
	MapPosition        int         `json:"mapPosition"`
	OpponentAttacks    int         `json:"opponentAttacks"`
}

type WarAttack struct {
	AttackerTag           string `json:"attackerTag"`
	DefenderTag           string `json:"defenderTag"`
	Stars                 int    `json:"stars"`
	DestructionPercentage int    `json:"destructionPercentage"`
	Order                 int    `json:"order"`
	Duration              int    `json:"duration"`
}

type Result string

const (
	Lose Result = "lose"
	Tie  Result = "tie"
	Win  Result = "win"
)

// Returns the War PerparationStart as a time.Time object
func (c *CurrentWar) PreparationStartTime() time.Time {
	parsed, _ := time.Parse(TimestampFormat, c.PreparationStart)
	return parsed
}

// Returns the War Start as a time.Time object
func (c *CurrentWar) StartTime() time.Time {
	parsed, _ := time.Parse(TimestampFormat, c.Start)
	return parsed
}

// Returns the War End as a time.Time object
func (c *CurrentWar) EndTime() time.Time {
	parsed, _ := time.Parse(TimestampFormat, c.End)
	return parsed
}

// Returns the War Start as a time.Time object
func (w *War) StartTime() time.Time {
	parsed, _ := time.Parse(TimestampFormat, w.Start)
	return parsed
}

// Returns the War End as a time.Time object
func (w *War) EndTime() time.Time {
	parsed, _ := time.Parse(TimestampFormat, w.End)
	return parsed
}

func (w *War) Won() bool {
	return w.Result == Win
}

func (w *War) Lost() bool {
	return w.Result == Lose
}

func (w *War) Tied() bool {
	return w.Result == Tie
}
