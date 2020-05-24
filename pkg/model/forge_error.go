package model

type ForgeError struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}
