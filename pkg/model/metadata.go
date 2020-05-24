package model

type Metadata struct {
	Name                   string                   `json:"name"`
	Version                string                   `json:"version"`
	Author                 string                   `json:"author"`
	Summary                string                   `json:"summary"`
	License                string                   `json:"license"`
	Source                 string                   `json:"source"`
	ProjectPage            string                   `json:"project_page"`
	IssuesUrl              string                   `json:"issues_url"`
	Dependencies           []Dependency             `json:"dependencies"`
	DataProvider           interface{}              `json:"data_provider"`
	Description            string                   `json:"description"`
	Tags                   []string                 `json:"tags"`
	Requirements           []Dependency             `json:"requirements"`
	OperatingsystemSupport []OperatingsystemSupport `json:"operatingsystem_support"`
}
