package model

type ModuleInfo struct {
	Slug
	Name         string `json:"name"`
	DeprecatedAt string `json:"deprecated_at"`
	Owner        Owner  `json:"owner"`
}
