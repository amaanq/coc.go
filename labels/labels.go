package labels

type PlayerLabelID int

func (p PlayerLabelID) Valid() bool {
	return p >= PlayerLabelClanWars && p <= PlayerLabelAmateurAttacker
}

type ClanLabelID int

func (c ClanLabelID) Valid() bool {
	return c >= ClanLabelClanWars && c <= ClanLabelNewbieFriendly
}

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
	IconUrls IconUrls `json:"iconUrls,omitempty"`
}

type IconUrls struct {
	Small  string `json:"small,omitempty"`
	Medium string `json:"medium,omitempty"`
}

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
