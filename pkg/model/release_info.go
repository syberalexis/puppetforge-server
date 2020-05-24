package model

import "time"

type ReleaseInfo struct {
	Slug
	Version   string    `json:"version"`
	Supported bool      `json:"supported"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"delete_at"`
	FileUri   string    `json:"file_uri"`
	FileSize  int       `json:"file_size"`
}
