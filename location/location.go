package location

type LocationData struct {
	Locations []Location `json:"items,omitempty"`
	Paging    struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type Location struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	IsCountry   bool   `json:"isCountry,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

type BadgeUrls struct {
	Small  string `json:"small,omitempty"`
	Medium string `json:"medium,omitempty"`
	Large  string `json:"large,omitempty"`
}
