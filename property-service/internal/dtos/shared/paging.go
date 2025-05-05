package dtos

type FilterModel struct {
	Type   string      `json:"type"`
	Filter interface{} `json:"filter"`
}
type SortModelItem struct {
	ColId string `json:"colId"`
	Sort  string `json:"sort"`
}
type SearchAdvanceModel struct {
	Filters  map[string]FilterModel `json:"filters"`
	Sort     []SortModelItem
	StartRow int `json:"startRow"`
	EndRow   int `json:"endRow"`
}

type SearchAdvanceResponse[T any] struct {
	Rows  []T
	Total int
}
