package model

type OperatingsystemSupport struct {
	Operatingsystem        string   `json:"operatingsystem"`
	Operatingsystemrelease []string `json:"operatingsystemrelease"`
}
