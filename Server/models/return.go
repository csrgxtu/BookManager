package models

type Recs interface{}

type Results struct {
	Error bool
	Msg   string
	Total int
	Data  []Recs
}
