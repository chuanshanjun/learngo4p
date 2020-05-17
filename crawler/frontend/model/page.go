package model

type SearchResult struct {
	Hints    int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []interface{}
}
