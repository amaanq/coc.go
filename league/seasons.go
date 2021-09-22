package league

type SeasonData struct {
	Seasons []Season `json:"items,omitempty"`
	Paging  struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type Season struct {
	ID string `json:"id,omitempty"`
}

type SeasonInfo struct {
	Players []RankedPlayer `json:"items,omitempty"`
	Paging  struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type RankedPlayer struct {
	Tag         string `json:"tag,omitempty"`
	Name        string `json:"name,omitempty"`
	ExpLevel    int64  `json:"expLevel,omitempty"`
	Trophies    int64  `json:"trophies,omitempty"`
	AttackWins  int64  `json:"attackWins,omitempty"`
	DefenseWins int64  `json:"defenseWins,omitempty"`
	Rank        int64  `json:"rank,omitempty"`
	Clan        Clan   `json:"clan,omitempty"`
}

type Clan struct {
	Tag       string    `json:"tag,omitempty"`
	Name      string    `json:"name,omitempty"`
	BadgeUrls BadgeUrls `json:"badgeUrls,omitempty"`
}

type BadgeUrls struct {
	Small  string `json:"small,omitempty"`
	Large  string `json:"large,omitempty"`
	Medium string `json:"medium,omitempty"`
}
