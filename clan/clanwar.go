package clan

type WarLog struct {
	Wars   []War `json:"items,omitempty"`
	Paging struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type War struct {
	Result   Result `json:"result"`
	EndTime  string `json:"endTime,omitempty"`
	TeamSize int64  `json:"teamSize,omitempty"`
	Clan     Clan   `json:"clan,omitempty"`
	Opponent Clan   `json:"opponent,omitempty"`
}

type Result string
const (
	Lose Result = "lose"
	Tie  Result = "tie"
	Win  Result = "win"
)

type CurrentWar struct {
	State                string `json:"state,omitempty"`
	TeamSize             int64  `json:"teamSize,omitempty"`
	PreparationStartTime string `json:"preparationStartTime,omitempty"`
	StartTime            string `json:"startTime,omitempty"`
	EndTime              string `json:"endTime,omitempty"`
	Clan                 WarClan   `json:"clan,omitempty"`
	Opponent             WarClan   `json:"opponent,omitempty"`
}

type WarClan struct {
	Tag                   string      `json:"tag,omitempty"`
	Name                  string      `json:"name,omitempty"`
	BadgeUrls             BadgeUrls   `json:"badgeUrls,omitempty"`
	ClanLevel             int64       `json:"clanLevel,omitempty"`
	Attacks               int64       `json:"attacks,omitempty"`
	Stars                 int64       `json:"stars,omitempty"`
	DestructionPercentage float64     `json:"destructionPercentage,omitempty"`
	Members               []WarMember `json:"members,omitempty"`
}

type WarMember struct {
	Tag                string   `json:"tag,omitempty"`
	Name               string   `json:"name,omitempty"`
	TownhallLevel      int64    `json:"townhallLevel,omitempty"`
	MapPosition        int64    `json:"mapPosition,omitempty"`
	OpponentAttacks    int64    `json:"opponentAttacks,omitempty"`
	BestOpponentAttack Attack   `json:"bestOpponentAttack,omitempty"`
	Attacks            []Attack `json:"attacks,omitempty"`
}

type Attack struct {
	AttackerTag           *string `json:"attackerTag,omitempty"`
	DefenderTag           *string `json:"defenderTag,omitempty"`
	Stars                 *int64  `json:"stars,omitempty"`
	DestructionPercentage *int64  `json:"destructionPercentage,omitempty"`
	Order                 *int64  `json:"order,omitempty"`
	Duration              *int64  `json:"duration,omitempty"`
}
