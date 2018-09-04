package models

import "time"

//Tokens when the user receives the token from authorization server
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//AccessToken struct for access tokens
type AccessToken struct {
	Exp int64  `json:"exp"`
	Iat int64  `json:"iat"`
	Jti string `json:"jti"`
}

//RefreshToken struct for refresh tokens
type RefreshToken struct {
	UserUUID string
	Exp      time.Time
	Token    string
}
