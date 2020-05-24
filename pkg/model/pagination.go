package model

type Pagination struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	First    string `json:"first"`
	Previous string `json:"previous"`
	Current  string `json:"current"`
	Next     string `json:"next"`
	Total    int    `json:"total"`
}
