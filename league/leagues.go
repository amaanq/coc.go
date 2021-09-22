package league

type LeagueData struct {
	Leagues []League `json:"items,omitempty"`
	Paging  struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type League struct {
	ID       int64    `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	IconUrls IconUrls `json:"iconUrls,omitempty"`
}

type IconUrls struct {
	Small  string `json:"small,omitempty"`
	Tiny   string `json:"tiny,omitempty"`
	Medium string `json:"medium,omitempty"`
}
