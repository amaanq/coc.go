package coc

type LocationID int

func (l LocationID) Valid() bool {
	return l >= Europe && l <= Zimbabwe
}

type LocationData struct {
	Locations []Location `json:"items,omitempty"`
	Paging    Paging     `json:"paging,omitempty"`
}

type Location struct {
	LocalizedName string `json:"localizedName,omitempty"`
	ID            int  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	IsCountry     bool   `json:"isCountry,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
}
