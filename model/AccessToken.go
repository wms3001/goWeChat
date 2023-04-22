package model

type AccessToken struct {
	Base
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
