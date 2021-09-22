package location

type ClanData struct {
	Clans  []Clan `json:"items,omitempty"`
	Paging struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type ClanVersusData struct {
	ClansVersus []ClanVersus `json:"items,omitempty"`
	Paging      struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type Clan struct {
	Tag          string    `json:"tag,omitempty"`
	Name         string    `json:"name,omitempty"`
	Location     Location  `json:"location,omitempty"`
	BadgeUrls    BadgeUrls `json:"badgeUrls,omitempty"`
	ClanLevel    int64     `json:"clanLevel,omitempty"`
	Members      int64     `json:"members,omitempty"`
	ClanPoints   int64     `json:"clanPoints,omitempty"`
	Rank         int64     `json:"rank,omitempty"`
	PreviousRank int64     `json:"previousRank,omitempty"`
}

//same as clan except clanversuspoints has a different json field
type ClanVersus struct {
	Tag              string    `json:"tag,omitempty"`
	Name             string    `json:"name,omitempty"`
	Location         Location  `json:"location,omitempty"`
	BadgeUrls        BadgeUrls `json:"badgeUrls,omitempty"`
	ClanLevel        int64     `json:"clanLevel,omitempty"`
	Members          int64     `json:"members,omitempty"`
	ClanVersusPoints int64     `json:"clanVersusPoints,omitempty"`
	Rank             int64     `json:"rank,omitempty"`
	PreviousRank     int64     `json:"previousRank,omitempty"`
}
