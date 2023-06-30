package coc

type LocationID int

func (l LocationID) Valid() bool {
	return l >= Europe && l <= Zimbabwe
}

type LocationData struct {
	Paging    Paging     `json:"paging,omitempty"`
	Locations []Location `json:"items,omitempty"`
}

type Location struct {
	LocalizedName string `json:"localizedName,omitempty"`
	Name          string `json:"name,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
	ID            int    `json:"id,omitempty"`
	IsCountry     bool   `json:"isCountry,omitempty"`
}
