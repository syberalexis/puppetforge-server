package model

import "time"

type Release struct {
	ReleaseInfo
	Module          ModuleInfo  `json:"module"`
	Metadata        Metadata    `json:"metadata"`
	Tags            []string    `json:"tags"`
	Pdk             bool        `json:"pdk"`
	ValidationScore int         `json:"validation_score"`
	FileMd5         string      `json:"file_md5"`
	FileSha256      string      `json:"file_sha256"`
	Downloads       int         `json:"downloads"`
	Readme          string      `json:"readme"`
	Changelog       string      `json:"changelog"`
	License         string      `json:"license"`
	Reference       string      `json:"reference"`
	Docs            interface{} `json:"docs"`
	Tasks           []Task      `json:"tasks"`
	Plans           []Plan      `json:"plans"`
	UpdatedAt       time.Time   `json:"updated_at"`
	DeletedFor      interface{} `json:"deleted_for"`
}
