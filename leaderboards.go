package coc

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// /rankings/clans
//_______________________________________________________________________

type ClanRankingList struct {
	Paging Paging        `json:"paging,omitempty"`
	Clans  []ClanRanking `json:"items,omitempty"`
}

type ClanRanking struct {
	BadgeURLs    IconURLs `json:"badgeUrls,omitempty"`
	Tag          string   `json:"tag,omitempty"`
	Name         string   `json:"name,omitempty"`
	Location     Location `json:"location,omitempty"`
	ClanPoints   int      `json:"clanPoints,omitempty"`
	ClanLevel    int      `json:"clanLevel,omitempty"`
	Rank         int      `json:"rank,omitempty"`
	PreviousRank int      `json:"previousRank,omitempty"`
	Members      int      `json:"members,omitempty"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// /rankings/players
//_______________________________________________________________________

type PlayerRankingList struct {
	Paging  Paging          `json:"paging,omitempty"`
	Players []PlayerRanking `json:"items,omitempty"`
}

type PlayerRanking struct {
	Clan         PlayerRankingClan `json:"clan,omitempty"`
	Tag          string            `json:"tag,omitempty"`
	Name         string            `json:"name,omitempty"`
	League       League            `json:"league,omitempty"`
	ExpLevel     int               `json:"expLevel,omitempty"`
	Rank         int               `json:"rank,omitempty"`
	PreviousRank int               `json:"previousRank,omitempty"`
	Trophies     int               `json:"trophies,omitempty"`
	AttackWins   int               `json:"attackWins,omitempty"`
	DefenseWins  int               `json:"defenseWins,omitempty"`
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
	Paging      Paging              `json:"paging,omitempty"`
	ClansVersus []ClanVersusRanking `json:"items,omitempty"`
}

type ClanVersusRanking struct {
	BadgeURLs        IconURLs `json:"badgeUrls,omitempty"`
	Tag              string   `json:"tag,omitempty"`
	Name             string   `json:"name,omitempty"`
	Location         Location `json:"location,omitempty"`
	ClanVersusPoints int      `json:"clanVersusPoints,omitempty"`
	ClanLevel        int      `json:"clanLevel,omitempty"`
	Rank             int      `json:"rank,omitempty"`
	PreviousRank     int      `json:"previousRank,omitempty"`
	Members          int      `json:"members,omitempty"`
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// /rankings/players-versus
//_______________________________________________________________________

type PlayerVersusRankingList struct {
	Paging        Paging                `json:"paging,omitempty"`
	PlayersVersus []PlayerVersusRanking `json:"items,omitempty"`
}

type PlayerVersusRanking struct {
	Clan           PlayerRankingClan `json:"clan,omitempty"`
	Tag            string            `json:"tag,omitempty"`
	Name           string            `json:"name,omitempty"`
	ExpLevel       int               `json:"expLevel,omitempty"`
	Rank           int               `json:"rank,omitempty"`
	PreviousRank   int               `json:"previousRank,omitempty"`
	VersusTrophies int               `json:"versusTrophies,omitempty"`
}
