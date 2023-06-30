package coc

type SeasonData struct {
	Paging  Paging   `json:"paging,omitempty"`
	Seasons []Season `json:"items,omitempty"`
}

type SeasonInfo struct {
	Paging  Paging         `json:"paging,omitempty"`
	Players []RankedPlayer `json:"items,omitempty"`
}

type RankedPlayer struct {
	Tag         string `json:"tag,omitempty"`
	Name        string `json:"name,omitempty"`
	Clan        Clan   `json:"clan,omitempty"`
	ExpLevel    int    `json:"expLevel,omitempty"`
	Trophies    int    `json:"trophies,omitempty"`
	AttackWins  int    `json:"attackWins,omitempty"`
	DefenseWins int    `json:"defenseWins,omitempty"`
	Rank        int    `json:"rank,omitempty"`
}
