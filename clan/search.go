package clan

type ClanList struct {
	Clans  []Clan `json:"items,omitempty"`
	Paging struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}
