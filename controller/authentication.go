package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bhupeshbhatia/go-authserver/models"
	"github.com/bhupeshbhatia/go-authserver/service"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

//Tokens contains access and refresh token
type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type payload struct {
	newToken string
	body     string
}

//ValidateAccessToken checks if the token exists and whether it is valid
func ValidateAccessToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := service.ParseAndDecryptToken(r)
	if err != nil {
		err = errors.Wrap(err, "Error parsing JWT")
		log.Println(err)
	}

	//Check if access token is valid
	if service.IsAccessTokenValid(token) {
		next(w, r)
	} else {
		//If access token is not valid then check newAccessToken func is called
		token := newAccessToken(token) //Should this be moved to jwt.go (service)?

		if token == "New refresh token required!" {
			w.Write([]byte("Need refresh Token"))
			http.HandleFunc("/login", Login)
		} else {
			next(w, r)
		}
	}
}

//newAccessToken is used when previous access token has expired. It calls IsRefreshTokenValid to check if new access token can be generated.
func newAccessToken(token *jwt.Token) string {
	if service.IsRefreshTokenValid(token) {
		userUUID := service.GetUUIDFromRedis(token)
		token, err := service.GenerateAccessToken(userUUID)
		if err != nil {
			err = errors.Wrap(err, "Access token not generated")
			log.Println(err)
		}
		return token
	}
	return "New refresh token required!"
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

	fmt.Println(service.AuthenticateUser(requestUser))

	//Authenticate user using information from request body
	if !service.AuthenticateUser(requestUser) {
		w.Write([]byte("Unable to locate user. Please sign in"))
	}

	accessToken, err := service.GenerateAccessToken(requestUser.UUID)
	if err != nil {
		err = errors.Wrap(err, "Access token not generated.")
		log.Println(err)
	}

	refreshToken, err := service.GenerateRefreshToken()
	if err != nil {
		err = errors.Wrap(err, "Refresh token not generated.")
		log.Println(err)
	}

	tokens := Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	accessAndRefreshTokens, err := json.Marshal(tokens)
	if err != nil {
		err = errors.Wrap(err, "token json not created.")
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(accessAndRefreshTokens)
}

//FileInsideServer for testing
func FileInsideServer(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	w.Write([]byte("This is the file I am trying to access"))
}

//HelloController for testing
func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello controller"))
}
