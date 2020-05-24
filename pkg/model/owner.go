package model

type Owner struct {
	Slug
	Username   string `json:"username"`
	GravatarId string `json:"gravatar_id"`
}
