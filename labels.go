package coc

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Player
//_______________________________________________________________________

type PlayerLabelID int

func (p PlayerLabelID) Valid() bool {
	return p >= PlayerLabelClanWars && p <= PlayerLabelAmateurAttacker
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Clan
//_______________________________________________________________________

type ClanLabelID int

func (c ClanLabelID) Valid() bool {
	return c >= ClanLabelClanWars && c <= ClanLabelNewbieFriendly
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Generic Label
//_______________________________________________________________________

type LabelsData struct {
	Labels []Label `json:"items,omitempty"`
	Paging struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type Label struct {
	ID       int64    `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	IconUrls IconURLs `json:"iconUrls,omitempty"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Constants
//_______________________________________________________________________

const (
	ClanLabelClanWars ClanLabelID = 56000000 + iota
	ClanLabelClanWarLeague
	ClanLabelTrophyPushing
	ClanLabelFriendlyWars
	ClanLabelClanGames
	ClanLabelBuilderBase
	ClanLabelBaseDesigning
	ClanLabelInternational
	ClanLabelFarming
	ClanLabelDonations
	ClanLabelFriendly
	ClanLabelTalkative
	ClanLabelUnderdog
	ClanLabelRelaxed
	ClanLabelCompetitive
	ClanLabelNewbieFriendly

	PlayerLabelClanWars PlayerLabelID = 57000000 + iota - 16
	PlayerLabelClanWarLeague
	PlayerLabelTrophyPushing
	PlayerLabelFriendlyWars
	PlayerLabelClanGames
	PlayerLabelBuilderBase
	PlayerLabelBaseDesigning
	PlayerLabelFarming
	PlayerLabelActiveDonator
	PlayerLabelActiveDaily
	PlayerLabelHungryLearner
	PlayerLabelFriendly
	PlayerLabelTalkative
	PlayerLabelTeacher
	PlayerLabelCompetitive
	PlayerLabelVeteran
	PlayerLabelNewbie
	PlayerLabelAmateurAttacker
)
