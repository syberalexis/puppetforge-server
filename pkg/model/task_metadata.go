package model

type TaskMetadata struct {
	Parameters  map[string]TaskParameter `json:"parameters"`
	Description string                   `json:"description"`
	InputMethod string                   `json:"input_method"`
}
