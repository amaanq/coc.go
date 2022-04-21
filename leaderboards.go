package coc

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// /rankings/clans
//_______________________________________________________________________

type ClanRankingList struct {
	Clans  []ClanRanking `json:"items,omitempty"`
	Paging Paging        `json:"paging,omitempty"`
}

type ClanRanking struct {
	Tag          string   `json:"tag,omitempty"`
	Name         string   `json:"name,omitempty"`
	ClanPoints   int      `json:"clanPoints,omitempty"`
	ClanLevel    int      `json:"clanLevel,omitempty"`
	Location     Location `json:"location,omitempty"`
	Rank         int      `json:"rank,omitempty"`
	PreviousRank int      `json:"previousRank,omitempty"`
	BadgeURLs    IconURLs `json:"badgeUrls,omitempty"`
	Members      int      `json:"members,omitempty"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// /rankings/players
//_______________________________________________________________________

type PlayerRankingList struct {
	Players []PlayerRanking `json:"items,omitempty"`
	Paging  Paging          `json:"paging,omitempty"`
}

type PlayerRanking struct {
	Tag          string            `json:"tag,omitempty"`
	Name         string            `json:"name,omitempty"`
	ExpLevel     int               `json:"expLevel,omitempty"`
	Rank         int               `json:"rank,omitempty"`
	PreviousRank int               `json:"previousRank,omitempty"`
	Trophies     int               `json:"trophies,omitempty"`
	AttackWins   int               `json:"attackWins,omitempty"`
	DefenseWins  int               `json:"defenseWins,omitempty"`
	Clan         PlayerRankingClan `json:"clan,omitempty"`
	League       League            `json:"league,omitempty"`
}

type PlayerRankingClan struct {
	Tag       string   `json:"tag,omitempty"`
	Name      string   `json:"name,omitempty"`
	BadgeURLs IconURLs `json:"badgeUrls,omitempty"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// /rankings/clans-versus
//_______________________________________________________________________

type ClanVersusRankingList struct {
	ClansVersus []ClanVersusRanking `json:"items,omitempty"`
	Paging      Paging              `json:"paging,omitempty"`
}

type ClanVersusRanking struct {
	Tag              string   `json:"tag,omitempty"`
	Name             string   `json:"name,omitempty"`
	ClanVersusPoints int      `json:"clanVersusPoints,omitempty"`
	ClanLevel        int      `json:"clanLevel,omitempty"`
	Location         Location `json:"location,omitempty"`
	Rank             int      `json:"rank,omitempty"`
	PreviousRank     int      `json:"previousRank,omitempty"`
	BadgeURLs        IconURLs `json:"badgeUrls,omitempty"`
	Members          int      `json:"members,omitempty"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// /rankings/players-versus
//_______________________________________________________________________

type PlayerVersusRankingList struct {
	PlayersVersus []PlayerVersusRanking `json:"items,omitempty"`
	Paging        Paging                `json:"paging,omitempty"`
}

type PlayerVersusRanking struct {
	Tag            string            `json:"tag,omitempty"`
	Name           string            `json:"name,omitempty"`
	ExpLevel       int               `json:"expLevel,omitempty"`
	Rank           int               `json:"rank,omitempty"`
	PreviousRank   int               `json:"previousRank,omitempty"`
	VersusTrophies int               `json:"versusTrophies,omitempty"`
	Clan           PlayerRankingClan `json:"clan,omitempty"`
}
