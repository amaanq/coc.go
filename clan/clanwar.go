package clan

import "time"

type WarLog struct {
	Items []War `json:"items"`
}

type War struct {
	Result           Result  `json:"result"`
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
	State                string  `json:"state"`
	TeamSize             int     `json:"teamSize"`
	AttacksPerMember     int     `json:"attacksPerMember"`
	PreparationStartTime string  `json:"preparationStartTime"`
	StartTime            string  `json:"startTime"`
	EndTime              string  `json:"endTime"`
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

func (w *War) EndTime() time.Time {
	parsed, _ := time.Parse("20060102T150405.999Z", w.endTime)
	return parsed
}
