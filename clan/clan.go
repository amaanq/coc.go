package clan

type WarFrequency string

const (
	Always       WarFrequency = "always"
	TwicePerWeek WarFrequency = "moreThanOncePerWeek"
	OncePerWeek  WarFrequency = "oncePerWeek"
	Rarely       WarFrequency = "lessThanOncePerWeek"
	Never        WarFrequency = "never"
	NotSet       WarFrequency = "unknown" // Doesn't seem to work..
)

func (w WarFrequency) Valid() bool {
	switch w {
	case Always, TwicePerWeek, OncePerWeek, Rarely, Never, NotSet:
		return true
	}
	return false
}

type Clan struct {
	Tag                    string       `json:"tag"`
	Name                   string       `json:"name"`
	Type                   string       `json:"type"`
	Description            string       `json:"description"`
	Location               Location     `json:"location"`
	BadgeUrls              BadgeUrls    `json:"badgeUrls"`
	ClanLevel              int          `json:"clanLevel"`
	ClanPoints             int          `json:"clanPoints"`
	ClanVersusPoints       int          `json:"clanVersusPoints"`
	RequiredTrophies       int          `json:"requiredTrophies"`
	WarFrequency           WarFrequency `json:"warFrequency"`
	WarWinStreak           int          `json:"warWinStreak"`
	WarWins                int          `json:"warWins"`
	WarTies                int          `json:"warTies"`
	WarLosses              int          `json:"warLosses"`
	IsWarLogPublic         bool         `json:"isWarLogPublic"`
	WarLeague              WarLeague    `json:"warLeague"`
	Members                int          `json:"members"`
	MemberList             []ClanMember `json:"memberList"`
	Labels                 []Label      `json:"labels"`
	ChatLanguage           ChatLanguage `json:"chatLanguage"`
	RequiredVersusTrophies int          `json:"requiredVersusTrophies"`
	RequiredTownhallLevel  int          `json:"requiredTownhallLevel"`
}

type ClanMembers struct {
	// We don't care about paging since a clan members list is not that large
	ClanMembers []ClanMember `json:"items"`
}

type ClanMember struct {
	Tag               string `json:"tag"`
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

type BadgeUrls struct {
	Small  string `json:"small"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type ChatLanguage struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	LanguageCode string `json:"languageCode"`
}

type Label struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	IconUrls LabelIconUrls `json:"iconUrls"`
}

type LabelIconUrls struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
}

type Location struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsCountry   bool   `json:"isCountry"`
	CountryCode string `json:"countryCode"`
}

type League struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	IconUrls LeagueIconUrls `json:"iconUrls"`
}

type LeagueIconUrls struct {
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
	Medium string `json:"medium"`
}

type WarLeague struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Role string

const (
	Elder    Role = "admin"
	CoLeader Role = "coLeader"
	Leader   Role = "leader"
	Member   Role = "member"
)

func (r Role) String() string {
	switch r {
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
