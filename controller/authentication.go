package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bhupeshbhatia/go-authserver/models"
	"github.com/bhupeshbhatia/go-authserver/service"
	"github.com/pkg/errors"
)

//Tokens contains access and refresh token
type Tokens struct {
	AccessToken  string
	RefreshToken string
}

//ValidateAccessToken checks if the token exists and whether it is valid
func ValidateAccessToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := service.ParseAndDecryptToken(r)

	if err != nil {
		err = errors.Wrap(err, "Error parsing JWT")
		log.Println(err)
	}

	if service.IsAccessTokenValid(token) {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//Login user by reading the payload, authenticating user, creating access and refresh tokens
func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)

	//Get information from request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		err = errors.Wrap(err, "User cannot be decoded from request body.")
		log.Println(err)
	}

	//Authenticate user using information from request body
	if !service.AuthenticateUser(requestUser) {
		w.Write([]byte("Unable to locate user. Please sign in"))
	}

	accessToken, err := service.GenerateAccessToken(requestUser.UUID)
	if err != nil {
		err = errors.Wrap(err, "Access token not generated.")
		log.Println(err)
	}

	refreshToken, err := service.RefreshToken()
	if err != nil {
		err = errors.Wrap(err, "Refresh token not generated.")
		log.Println(err)
	}

	tokens := Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(tokens.AccessToken))
	w.Write([]byte(tokens.RefreshToken))
}

func FileInsideServer(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("This is the file I am trying to access"))
}

//HelloController for testing
func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello controller"))
}
