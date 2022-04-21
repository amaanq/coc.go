package coc

type WarFrequency string

const (
	NotSet       WarFrequency = "unknown" // Doesn't seem to work..
	Always       WarFrequency = "always"
	TwicePerWeek WarFrequency = "moreThanOncePerWeek"
	OncePerWeek  WarFrequency = "oncePerWeek"
	Rarely       WarFrequency = "lessThanOncePerWeek"
	Never        WarFrequency = "never"
	Any          WarFrequency = "any"
)

func (w WarFrequency) Valid() bool {
	switch w {
	case Always, TwicePerWeek, OncePerWeek, Rarely, Never, NotSet:
		return true
	}
	return false
}

type Clan struct {
	Tag                    ClanTag       `json:"tag,omitempty"`
	Name                   string       `json:"name,omitempty"`
	Type                   string       `json:"type,omitempty"`
	Description            string       `json:"description,omitempty"`
	Location               Location     `json:"location,omitempty"`
	BadgeURLs              IconURLs     `json:"badgeUrls,omitempty"`
	ClanLevel              int          `json:"clanLevel,omitempty"`
	ClanPoints             int          `json:"clanPoints,omitempty"`
	ClanVersusPoints       int          `json:"clanVersusPoints,omitempty"`
	RequiredTrophies       int          `json:"requiredTrophies,omitempty"`
	WarFrequency           WarFrequency `json:"warFrequency,omitempty"`
	WarWinStreak           int          `json:"warWinStreak,omitempty"`
	WarWins                int          `json:"warWins,omitempty"`
	WarTies                int          `json:"warTies,omitempty"`
	WarLosses              int          `json:"warLosses,omitempty"`
	IsWarLogPublic         bool         `json:"isWarLogPublic,omitempty"`
	WarLeague              WarLeague    `json:"warLeague,omitempty"`
	Members                int          `json:"members,omitempty"`
	MemberList             []ClanMember `json:"memberList,omitempty"`
	Labels                 []Label      `json:"labels,omitempty"`
	ChatLanguage           ChatLanguage `json:"chatLanguage,omitempty"`
	RequiredVersusTrophies int          `json:"requiredVersusTrophies,omitempty"`
	RequiredTownhallLevel  int          `json:"requiredTownhallLevel,omitempty"`
}

type ClanMembers struct {
	ClanMembers []ClanMember `json:"items,omitempty"`
	Paging      Paging       `json:"paging,omitempty"`
}

type ClanMember struct {
	Tag               PlayerTag `json:"tag"`
	Name              string `json:"name"`
	Role              Role   `json:"role"`
	ExpLevel          int    `json:"expLevel"`
	League            League `json:"league"`
	Trophies          int    `json:"trophies"`
	VersusTrophies    int    `json:"versusTrophies"`
	ClanRank          int    `json:"clanRank"`
	PreviousClanRank  int    `json:"previousClanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
}

type ChatLanguage struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	LanguageCode string `json:"languageCode"`
}

type LeagueIconUrls struct {
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
	Medium string `json:"medium"`
}

type Role string

const (
	NotMember Role = "notMember" // ?
	Elder     Role = "admin"
	CoLeader  Role = "coLeader"
	Leader    Role = "leader"
	Member    Role = "member"
)

func (r Role) String() string {
	switch r {
	case NotMember:
		return "Not A Member"
	case Elder:
		return "Elder"
	case CoLeader:
		return "CoLeader"
	case Leader:
		return "Leader"
	case Member:
		return "Member"
	}
	return ""
}
