package clan

type WarLog struct {
	Items []War `json:"items"`
}

type War struct {
	Result           Result `json:"result"`
	EndTime          string `json:"endTime"`
	TeamSize         int64  `json:"teamSize"`
	AttacksPerMember int64  `json:"attacksPerMember"`
	Clan             Clan   `json:"clan"`
	Opponent         Clan   `json:"opponent"`
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
	Tag             string `json:"tag"`
	Name            string `json:"name"`
	TownhallLevel   int64  `json:"townhallLevel"`
	MapPosition     int64  `json:"mapPosition"`
	OpponentAttacks int64  `json:"opponentAttacks"`
}

type Result string

const (
	Lose Result = "lose"
	Tie  Result = "tie"
	Win  Result = "win"
)
