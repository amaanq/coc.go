package coc

type ClanList struct {
	Paging Paging `json:"paging,omitempty"`
	Clans  []Clan `json:"items,omitempty"`
}
