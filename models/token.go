package models

//TokenAuthentication when the user receives the token from authorization server
type TokenAuthentication struct {
	Token string `json:"token"`
}
