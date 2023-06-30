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
	Paging Paging  `json:"paging,omitempty"`
	Labels []Label `json:"items,omitempty"`
}

type Label struct {
	IconUrls IconURLs `json:"iconUrls,omitempty"`
	Name     string   `json:"name,omitempty"`
	ID       int      `json:"id,omitempty"`
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
