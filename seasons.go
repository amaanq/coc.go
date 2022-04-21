package coc

type SeasonData struct {
	Seasons []Season `json:"items,omitempty"`
	Paging  struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type SeasonInfo struct {
	Players []RankedPlayer `json:"items,omitempty"`
	Paging  struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type RankedPlayer struct {
	Tag         PlayerTag `json:"tag,omitempty"`
	Name        string    `json:"name,omitempty"`
	ExpLevel    int64     `json:"expLevel,omitempty"`
	Trophies    int64     `json:"trophies,omitempty"`
	AttackWins  int64     `json:"attackWins,omitempty"`
	DefenseWins int64     `json:"defenseWins,omitempty"`
	Rank        int64     `json:"rank,omitempty"`
	Clan        Clan      `json:"clan,omitempty"`
}
