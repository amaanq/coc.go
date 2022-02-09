package clan

type WarLog struct {
	Items []War `json:"items"`
}

type War struct {
	Result           Result  `json:"result"`
	EndTime          string  `json:"endTime"`
	TeamSize         int64   `json:"teamSize"`
	AttacksPerMember int64   `json:"attacksPerMember"`
	Clan             WarClan `json:"clan"`
	Opponent         WarClan `json:"opponent"`
}

type WarClan struct {
	Tag                   string      `json:"tag"`
	Name                  string      `json:"name"`
	BadgeUrls             BadgeUrls   `json:"badgeUrls"`
	ClanLevel             int64       `json:"clanLevel"`
	Attacks               int64       `json:"attacks"`
	Stars                 int64       `json:"stars"`
	DestructionPercentage float64     `json:"destructionPercentage"`
	ExpEarned             int64       `json:"expEarned"`
	Members               []WarMember `json:"members"`
}

type CurrentWar struct {
	State                string  `json:"state"`
	TeamSize             int64   `json:"teamSize"`
	AttacksPerMember     int64   `json:"attacksPerMember"`
	PreparationStartTime string  `json:"preparationStartTime"`
	StartTime            string  `json:"startTime"`
	EndTime              string  `json:"endTime"`
	Clan                 WarClan `json:"clan"`
	Opponent             WarClan `json:"opponent"`
}

type WarMember struct {
	Tag                string      `json:"tag"`
	Name               string      `json:"name"`
	TownhallLevel      int64       `json:"townhallLevel"`
	MapPosition        int64       `json:"mapPosition"`
	Attacks            []WarAttack `json:"attacks"`
	OpponentAttacks    int64       `json:"opponentAttacks"`
	BestOpponentAttack WarAttack   `json:"bestOpponentAttack"`
}

type WarAttack struct {
	// Isn't attacker tag redundant lol
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
