package model

type Task struct {
	Name        string   `json:"name"`
	Executable  string   `json:"executable"`
	Executables []string `json:"executables"`
	Description string   `json:"description"`
	Metadata    Metadata `json:"metadata"`
}
