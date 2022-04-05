package warleague

type WarLeagueID int

func (w WarLeagueID) Valid() bool {
	return w >= Unranked && w <= ChampionLeagueI
}

const (
	Unranked WarLeagueID = 48000000 + iota
	BronzeLeagueIII
	BronzeLeagueII
	BronzeLeagueI
	SilverLeagueIII
	SilverLeagueII
	SilverLeagueI
	GoldLeagueIII
	GoldLeagueII
	GoldLeagueI
	CrystalLeagueIII
	CrystalLeagueII
	CrystalLeagueI
	MasterLeagueIII
	MasterLeagueII
	MasterLeagueI
	ChampionLeagueIII
	ChampionLeagueII
	ChampionLeagueI
)

type LeagueData struct {
	Leagues []League `json:"items,omitempty"`
	Paging  struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type League struct {
	ID       int64    `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	IconUrls IconUrls `json:"iconUrls,omitempty"`
}

type IconUrls struct {
	Small  string `json:"small,omitempty"`
	Tiny   string `json:"tiny,omitempty"`
	Medium string `json:"medium,omitempty"`
}
