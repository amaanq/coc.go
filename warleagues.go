package coc

type WarLeagueID int

func (w WarLeagueID) Valid() bool {
	return w >= UnrankedWarLeague && w <= ChampionWarLeagueI
}

const (
	UnrankedWarLeague WarLeagueID = 48000000 + iota
	BronzeWarLeagueIII
	BronzeWarLeagueII
	BronzeWarLeagueI
	SilverWarLeagueIII
	SilverWarLeagueII
	SilverWarLeagueI
	GoldWarLeagueIII
	GoldWarLeagueII
	GoldWarLeagueI
	CrystalWarLeagueIII
	CrystalWarLeagueII
	CrystalWarLeagueI
	MasterWarLeagueIII
	MasterWarLeagueII
	MasterWarLeagueI
	ChampionWarLeagueIII
	ChampionWarLeagueII
	ChampionWarLeagueI
)

type LeagueData struct {
	Paging  Paging   `json:"paging,omitempty"`
	Leagues []League `json:"items,omitempty"`
}

type League struct {
	IconUrls IconURLs `json:"iconUrls,omitempty"`
	Name     string   `json:"name,omitempty"`
	ID       int      `json:"id,omitempty"`
}
