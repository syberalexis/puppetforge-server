package model

import "time"

type ModuleInfo struct {
	Slug
	Name         string    `json:"name"`
	DeprecatedAt time.Time `json:"deprecated_at"`
	Owner        Owner     `json:"owner"`
}
