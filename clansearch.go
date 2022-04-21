package coc

type ClanList struct {
	Clans  []Clan `json:"items,omitempty"`
	Paging Paging `json:"paging,omitempty"`
}
