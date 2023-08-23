package coc

import "fmt"

type Clan struct {
	MemberList             []ClanMember `json:"memberList"`
	WarTies                int          `json:"warTies"`
	Description            string       `json:"description"`
	Location               Location     `json:"location"`
	WarLosses              int          `json:"warLosses"`
	ChatLanguage           ChatLanguage `json:"chatLanguage"`
	ClanCapital            ClanCapital  `json:"clanCapital"`
	WarLeague              WarLeague    `json:"warLeague"`
	BadgeURLs              IconURLs     `json:"badgeUrls"`
	Tag                    string       `json:"tag"`
	WarFrequency           WarFrequency `json:"warFrequency"`
	Privacy                Privacy      `json:"type"`
	Name                   string       `json:"name"`
	Labels                 []Label      `json:"labels"`
	ClanLevel              int          `json:"clanLevel"`
	WarWins                int          `json:"warWins"`
	Members                int          `json:"members"`
	WarWinStreak           int          `json:"warWinStreak"`
	RequiredTrophies       int          `json:"requiredTrophies"`
	RequiredVersusTrophies int          `json:"requiredVersusTrophies"`
	RequiredTownhallLevel  int          `json:"requiredTownhallLevel"`
	ClanVersusPoints       int          `json:"clanVersusPoints"`
	ClanPoints             int          `json:"clanPoints"`
	IsWarLogPublic         bool         `json:"isWarLogPublic"`
}

func (c *Clan) GameLink() string {
	return fmt.Sprintf("https://link.clashofclans.com/en?action=OpenClanProfile&tag=%s", c.Tag[1:])
}

func (c *Clan) ClashOfStatsLink() string {
	return fmt.Sprintf("https://www.clashofstats.com/clans/%s/summary", c.Tag[1:])
}

func (c *Clan) ChocolateClashLink() string {
	return fmt.Sprintf("https://cc.chocolateclash.com/cc_n/clan.php?tag=%s", c.Tag[1:])
}

type Privacy string

const (
	Open       Privacy = "open"
	InviteOnly Privacy = "inviteOnly"
	Closed     Privacy = "closed"
)

func (p Privacy) String() string {
	switch p {
	case Open:
		return "Anyone Can Join"
	case InviteOnly:
		return "Invite Only"
	case Closed:
		return "Closed"
	}
	return ""
}

type WarFrequency string

const (
	Unknown             WarFrequency = "unknown" // Doesn't seem to work..
	Always              WarFrequency = "always"
	MoreThanOncePerWeek WarFrequency = "moreThanOncePerWeek"
	OncePerWeek         WarFrequency = "oncePerWeek"
	LessThanOncePerWeek WarFrequency = "lessThanOncePerWeek"
	Never               WarFrequency = "never"
	Any                 WarFrequency = "any"
)

func (w WarFrequency) IsUnknown() bool {
	return w == Unknown
}

func (w WarFrequency) IsAlways() bool {
	return w == Always
}

func (w WarFrequency) IsMoreThanOncePerWeek() bool {
	return w == MoreThanOncePerWeek
}

func (w WarFrequency) IsOncePerWeek() bool {
	return w == OncePerWeek
}

func (w WarFrequency) IsLessThanOncePerWeek() bool {
	return w == LessThanOncePerWeek
}

func (w WarFrequency) IsNever() bool {
	return w == Never
}

func (w WarFrequency) IsAny() bool {
	return w == Any
}

func (w WarFrequency) Valid() bool {
	return w == Unknown || w == Always || w == MoreThanOncePerWeek || w == OncePerWeek || w == LessThanOncePerWeek || w == Never || w == Any
}

func (p Privacy) IsOpen() bool {
	return p == Open
}

func (p Privacy) IsInviteOnly() bool {
	return p == InviteOnly
}

func (p Privacy) IsClosed() bool {
	return p == Closed
}

type ChatLanguage struct {
	Name         string `json:"name"`
	LanguageCode string `json:"languageCode"`
	ID           int    `json:"id"`
}

type ClanMembers struct {
	Paging      Paging       `json:"paging"`
	ClanMembers []ClanMember `json:"items"`
}

type ClanMember struct {
	Tag               string `json:"tag"`
	Name              string `json:"name"`
	Role              Role   `json:"role"`
	League            League `json:"league"`
	ExpLevel          int    `json:"expLevel"`
	Trophies          int    `json:"trophies"`
	VersusTrophies    int    `json:"versusTrophies"`
	ClanRank          int    `json:"clanRank"`
	PreviousClanRank  int    `json:"previousClanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
}

type Role string

const (
	NotMember Role = "notMember"
	Member    Role = "member"
	Elder     Role = "admin"
	CoLeader  Role = "coLeader"
	Leader    Role = "leader"
)

func (r Role) String() string {
	switch r {
	case NotMember:
		return "Not Member"
	case Member:
		return "Member"
	case Elder:
		return "Elder"
	case CoLeader:
		return "Co-Leader"
	case Leader:
		return "Leader"
	}
	return ""
}

func (r Role) IsNotMember() bool {
	return r == NotMember
}

func (r Role) IsMember() bool {
	return r == Member
}

func (r Role) IsElder() bool {
	return r == Elder
}

func (r Role) IsCoLeader() bool {
	return r == CoLeader
}

func (r Role) IsLeader() bool {
	return r == Leader
}

type ClanCapital struct {
	CapitalHallLevel *int        `json:"capitalHallLevel"`
	Districts        *[]District `json:"districts"`
}

type District struct {
	Name              string `json:"name"`
	ID                int    `json:"id"`
	DistrictHallLevel int    `json:"districtHallLevel"`
}

type LeagueIconUrls struct {
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
	Medium string `json:"medium"`
}
