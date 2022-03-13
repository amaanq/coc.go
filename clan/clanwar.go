package clan

import "time"

type WarLog struct {
	Log []War `json:"items"`
}

type War struct {
	Result           Result  `json:"result"`
	startTime        string  `json:"startTime"`
	endTime          string  `json:"endTime"`
	TeamSize         int     `json:"teamSize"`
	AttacksPerMember int     `json:"attacksPerMember"`
	Clan             WarClan `json:"clan"`
	Opponent         WarClan `json:"opponent"`
}

type WarClan struct {
	Tag                   string      `json:"tag"`
	Name                  string      `json:"name"`
	BadgeUrls             BadgeUrls   `json:"badgeUrls"`
	ClanLevel             int         `json:"clanLevel"`
	Attacks               int         `json:"attacks"`
	Stars                 int         `json:"stars"`
	DestructionPercentage float64     `json:"destructionPercentage"`
	ExpEarned             int         `json:"expEarned"`
	Members               []WarMember `json:"members"`
}

type CurrentWar struct {
	Result               Result  `json:"result"`
	State                string  `json:"state"`
	TeamSize             int     `json:"teamSize"`
	AttacksPerMember     int     `json:"attacksPerMember"`
	preparationStartTime string  `json:"preparationStartTime"`
	startTime            string  `json:"startTime"`
	endTime              string  `json:"endTime"`
	Clan                 WarClan `json:"clan"`
	Opponent             WarClan `json:"opponent"`
}

type WarMember struct {
	Tag                string      `json:"tag"`
	Name               string      `json:"name"`
	TownhallLevel      int         `json:"townhallLevel"`
	MapPosition        int         `json:"mapPosition"`
	Attacks            []WarAttack `json:"attacks"`
	OpponentAttacks    int         `json:"opponentAttacks"`
	BestOpponentAttack WarAttack   `json:"bestOpponentAttack"`
}

type WarAttack struct {
	// Isn't attacker tag redundant
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

func (c *CurrentWar) Won() bool {
	return c.Result == Win
}

func (c *CurrentWar) Lost() bool {
	return c.Result == Lose
}

func (c *CurrentWar) Tied() bool {
	return c.Result == Tie
}

func (c *CurrentWar) PreparationStartTime() time.Time {
	parsed, _ := time.Parse("20060102T150405.999Z", c.preparationStartTime)
	return parsed
}

func (c *CurrentWar) StartTime() time.Time {
	parsed, _ := time.Parse("20060102T150405.999Z", c.startTime)
	return parsed
}

func (c *CurrentWar) EndTime() time.Time {
	parsed, _ := time.Parse("20060102T150405.999Z", c.endTime)
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

func (w *War) StartTime() time.Time {
	parsed, _ := time.Parse("20060102T150405.999Z", w.startTime)
	return parsed
}

func (w *War) EndTime() time.Time {
	parsed, _ := time.Parse("20060102T150405.999Z", w.endTime)
	return parsed
}
