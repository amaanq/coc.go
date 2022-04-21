package coc

import "fmt"

type LeagueID int

func (l LeagueID) Valid() bool {
	return l >= Unranked && l <= LegendLeague
}

type WarLeagueData struct {
	Leagues []League `json:"items,omitempty"`
	Paging  Paging   `json:"paging,omitempty"`
}

type WarLeague struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	IconUrls IconURLs `json:"iconUrls,omitempty"`
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
	ErrInvalidLeague = fmt.Errorf("only Legends League is supported, to avoid this error pass in 29000022 (or LegendLeague) for the LeagueID argument")
)
