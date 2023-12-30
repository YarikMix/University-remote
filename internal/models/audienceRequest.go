package models

type AudienceRequest struct {
	AccessToken int64  `json:"access_token"`
	Audience    string `json:"audience"`
}
