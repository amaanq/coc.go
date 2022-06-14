package coc

type SeasonData struct {
	Seasons []Season `json:"items,omitempty"`
	Paging  Paging   `json:"paging,omitempty"`
}

type SeasonInfo struct {
	Players []RankedPlayer `json:"items,omitempty"`
	Paging  Paging         `json:"paging,omitempty"`
}

type RankedPlayer struct {
	Tag         PlayerTag `json:"tag,omitempty"`
	Name        string    `json:"name,omitempty"`
	ExpLevel    int       `json:"expLevel,omitempty"`
	Trophies    int       `json:"trophies,omitempty"`
	AttackWins  int       `json:"attackWins,omitempty"`
	DefenseWins int       `json:"defenseWins,omitempty"`
	Rank        int       `json:"rank,omitempty"`
	Clan        Clan      `json:"clan,omitempty"`
}
