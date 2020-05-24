package model

type Page struct {
	Pagination Pagination `json:"pagination"`
	Results    []Module   `json:"results"`
}
