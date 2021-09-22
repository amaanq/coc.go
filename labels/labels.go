package labels

type LabelsData struct {
	Labels []Label `json:"items,omitempty"`
	Paging struct {
		Cursors struct {
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type Label struct {
	ID       int64    `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	IconUrls IconUrls `json:"iconUrls,omitempty"`
}

type IconUrls struct {
	Small  string `json:"small,omitempty"`
	Medium string `json:"medium,omitempty"`
}
