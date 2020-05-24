package model

type ReleaseInfo struct {
	Slug
	Version   string `json:"version"`
	Supported bool   `json:"supported"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"delete_at"`
	FileUri   string `json:"file_uri"`
	FileSize  int    `json:"file_size"`
}
