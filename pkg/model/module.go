package model

type Module struct {
	ModuleInfo
	Downloads      int           `json:"downloads"`
	CreatedAt      string        `json:"created_at"`
	UpdatedAt      string        `json:"updated_at"`
	DeprecatedFor  interface{}   `json:"deprecated_for"`
	SupersededBy   interface{}   `json:"superseded_by"`
	Supported      bool          `json:"supported"`
	Endorsement    string        `json:"endorsement"`
	ModuleGroup    string        `json:"module_group"`
	CurrentRelease Release       `json:"current_release"`
	Releases       []ReleaseInfo `json:"releases"`
	FeedbackScore  int           `json:"feedback_score"`
	HomepageUrl    string        `json:"homepage_url"`
	IssuesUrl      string        `json:"issues_url"`
}

func (module *Module) Equals(m *Module) bool {
	if m != nil {
		return module.SlugName == m.SlugName
	}

	return false
}
