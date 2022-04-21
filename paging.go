package coc

type Paging struct {
	Cursors Cursors `json:"cursors,omitempty"`
}

type Cursors struct {
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}
