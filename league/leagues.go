package league

type LeagueID int

func (l LeagueID) Valid() bool {
	return l >= Unranked && l <= LegendLeague
}

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

const (
	Unranked LeagueID = 29000000 + iota
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
	TitanLeagueIII
	TitanLeagueII
	TitanLeagueI
	LegendLeague
)


var (
	ErrInvalidLeague = "Only Legends League is supported with this. Deferring to 29000022 (aka league.LegendLeague). To avoid this message being printed again, pass in 29000022 (or league.LegendLeague) for the LeagueID argument."
)
