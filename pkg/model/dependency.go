package model

type Dependency struct {
	Name               string `json:"name"`
	VersionRequirement string `json:"version_requirement"`
}
