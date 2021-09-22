package clan

type ClanMember struct {
	Tag      string `json:"tag,omitempty"`
	Name     string `json:"name,omitempty"`
	Role     Role   `json:"role,omitempty"`
	ExpLevel int64  `json:"expLevel,omitempty"`
	League   struct {
		ID       int64  `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		IconUrls struct {
			Small  *string `json:"small,omitempty"`
			Tiny   *string `json:"tiny,omitempty"`
			Medium *string `json:"medium,omitempty"`
		} `json:"iconUrls,omitempty"`
	} `json:"league,omitempty"`
	Trophies          int64 `json:"trophies,omitempty"`
	VersusTrophies    int64 `json:"versusTrophies,omitempty"`
	ClanRank          int64 `json:"clanRank,omitempty"`
	PreviousClanRank  int64 `json:"previousClanRank,omitempty"`
	Donations         int64 `json:"donations,omitempty"`
	DonationsReceived int64 `json:"donationsReceived,omitempty"`
}

type Clan struct {
	BadgeUrls    BadgeUrls `json:"badgeUrls"`
	ChatLanguage struct {
		ID           int64  `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
		LanguageCode string `json:"languageCode,omitempty"`
	} `json:"chatLanguage,omitempty"`
	ClanLevel        int64  `json:"clanLevel"`
	ClanPoints       int64  `json:"clanPoints"`
	ClanVersusPoints int64  `json:"clanVersusPoints"`
	Description      string `json:"description"`
	IsWarLogPublic   bool   `json:"isWarLogPublic"`
	Location         struct {
		CountryCode string `json:"countryCode"`
		ID          int64  `json:"id"`
		IsCountry   bool   `json:"isCountry"`
		Name        string `json:"name"`
	} `json:"location"`
	MemberList             []ClanMember `json:"memberList"`
	Members                int64        `json:"members"`
	Name                   string       `json:"name"`
	RequiredTrophies       int64        `json:"requiredTrophies"`
	RequiredVersusTrophies int64        `json:"requiredVersusTrophies,omitempty"`
	RequiredTownhallLevel  int64        `json:"requiredTownhallLevel,omitempty"`
	Tag                    string       `json:"tag"`
	Type                   string       `json:"type"`
	WarFrequency           string       `json:"warFrequency"`
	WarLosses              int64        `json:"warLosses"`
	WarTies                int64        `json:"warTies"`
	WarWinStreak           int64        `json:"warWinStreak"`
	WarWins                int64        `json:"warWins"`
	WarLeague              struct {
		ID   int64  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"warLeague,omitempty"`
	Labels []struct {
		ID       int64    `json:"id,omitempty"`
		Name     string   `json:"name,omitempty"`
		IconUrls IconUrls `json:"iconUrls,omitempty"`
	} `json:"labels,omitempty"`
}

type Role string
const (
	Admin    Role = "admin"
	CoLeader Role = "coLeader"
	Leader   Role = "leader"
	Member   Role = "member"
)

type MemberList struct {
	Members []ClanMember `json:"items,omitempty"`
	Paging  struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type IconUrls struct {
	Small  string `json:"small,omitempty"`
	Tiny   string `json:"tiny,omitempty"`
	Medium string `json:"medium,omitempty"`
}

type BadgeUrls struct {
	Small  *string `json:"small,omitempty"`
	Large  *string `json:"large,omitempty"`
	Medium *string `json:"medium,omitempty"`
}
