package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhupeshbhatia/go-authserver/models"
	"github.com/bhupeshbhatia/go-authserver/service"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

var EndPoint interface {
	Login(w http.ResponseWriter, r *http.Request)
}

var EndPointAccess struct {
	ep *EndPoint
}

//Tokens contains access and refresh token
//Login user by reading the payload, authenticating user, creating access and refresh tokens
func (e *EndPointAccess) Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)

	//Get information from request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		err = errors.Wrap(err, "User cannot be decoded from request body.")
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	}

	user, err := service.AuthenticateUser(requestUser)
	if err != nil {
		log.Println("User not found during authentication")
		w.WriteHeader(http.StatusUnauthorized)
	}

	if user != nil {
		accessToken, err := service.GenerateAccessToken(user.UUID)
		if err != nil {
			err = errors.Wrap(err, "Access token not generated.")
			log.Println(err)
			return
		}

		refreshToken, err := service.GenerateRefreshToken(user.UUID)
		if err != nil {
			err = errors.Wrap(err, "Refresh token not generated.")
			log.Println(err)
			return
		}

		//THIS IS WHERE THE PROBLEM STARTS
		// exp, err := time.ParseDuration("168h")
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		service.RefreshTokens["refreshToken"] = models.RefreshToken{
			UserUUID: user.UUID,
			Exp:      time.Now().AddDate(0, 0, 7),
		}

		tokens := getToken(accessToken, refreshToken)
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("+++++++++++++++++++++++++++++")
		fmt.Println(tokens)
		w.Write(tokens)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func getToken(accessToken string, refreshToken string) []byte {
	tokens := models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	accessAndRefreshTokens, err := json.Marshal(tokens)
	if err != nil {
		err = errors.Wrap(err, "token json not created.")
		log.Println(err)
	}
	return accessAndRefreshTokens
}

//RefreshToken creates a refresh token for jwt
func RefreshToken(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// /*
	generatedUUID, err := uuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "UUID not generated.")
		log.Println(err)
	}

	//Convert UUID to byte[]
	convertedByteArrayUUID, err := generatedUUID.MarshalText()
	if err != nil {
		err = errors.Wrap(err, "UUID not converted to byte[].")
		log.Println(err)
	}

	//Changed byte[] to string
	uniqueID := string(convertedByteArrayUUID)

	// */

	refreshToken, err := service.GenerateRefreshToken(uniqueID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	w.Write([]byte(refreshToken))
}

//AccessToken gets a new access token using RefreshToken
func AccessToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println("=====================")
	fmt.Println(r.Header)

	// parsedToken, err := service.ParseAndDecryptToken(r)
	// if err != nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// expiry := parsedToken.Claims.(jwt.MapClaims)["exp"]
	// userUUID := parsedToken.Claims.(jwt.MapClaims)["jti"]
	// fmt.Println(expiry)
	// fmt.Println(userUUID)

	// fmt.Println(token)

	//From Db get expiry time of refresh token

	// refreshToken := service.RefreshToken {
	// 	UserUUID:
	// }

	// test, ok := service.RefreshTokens["refreshToken"]
	// fmt.Println(test)
	// fmt.Println(ok)
	// tokenStruct := service.RefreshTokens["refreshToken"]
	// expiryTime := tokenStruct.Exp
	// refreshToken := service.RefreshTokens[token] //this is wrong. there are no claims in refresh token. Need to check with db -- that's why refreshToken.Exp is empty
	// fmt.Println(refreshToken)

	// fmt.Println(service.IsTokenValid(refreshToken.Exp, time.Hour*168)) //false

	// fmt.Println(expiryTime)
	expiryTime := service.RefreshTokens["refreshToken"].Exp
	fmt.Println(service.RefreshTokens["refreshToken"].Exp) //refresh is negative

	//time is higher
	// See which value is higher/lower?
	// This is the issue them.
	// How about this guy^
	// if !service.IsTokenValid(refreshToken.Exp, time.Hour*168) {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	if !service.IsTokenValid(expiryTime, time.Hour*168) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("This is where it failed"))
		return
	}

	// accessToken, err := service.GenerateAccessToken(refreshToken.UserUUID)

	//Are we not trying to get this information from db? --
	accessToken, err := service.GenerateAccessToken(service.RefreshTokens["refreshToken"].UserUUID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(accessToken))
}

// //ValidateTokens checks if the token exists and whether it is valid
// func ValidateAccessToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	token, err := service.ParseAndDecryptToken(r)
// 	if err != nil {
// 		err = errors.Wrap(err, "Error parsing JWT")
// 		log.Println(err)
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	claims := token.Claims.(jwt.MapClaims)

// 	expiryTime := claims["exp"].(float64)
// 	// fmt.Println(claims["exp"])

// 	fmt.Println(expiryTime, float64(time.Now().Unix()))

// 	addDuration := time.Second * 0

// 	fmt.Println(service.IsTokenValid(expiryTime, addDuration))

// 	if service.IsTokenValid(expiryTime, addDuration) {
// 		next(w, r)
// 	} else {
// 		w.WriteHeader(http.StatusUnauthorized)
// 	}
// }

// //CheckHeaderAuthorization check if header has token
// func CheckHeaderAuthorization(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	authHeader := r.Header.Get("Authorization")

// 	if authHeader == "" {
// 		fmt.Println("auth empty")
// 		http.Error(w, "Forbidden", http.StatusForbidden)
// 	} else {
// 		fmt.Println("auth header not empty")
// 		next(w, r)
// 	}
// }

//ValidateTokens checks if the token exists and whether it is valid
// func ValidateTokens(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	token, err := service.ParseAndDecryptToken(r)
// 	if err != nil {
// 		err = errors.Wrap(err, "Error parsing JWT")
// 		log.Println(err)
// 	}

// 	//Generating UUID
// 	// userUUID := service.GetUUIDFromAccessToken(token)

// 	//Getting refresh token from Database using UUID
// 	refreshTokenFromRedis := service.GetRefreshTokenFromRedis(userUUID)

// 	//Check if access token is valid
// 	// if service.IsAccessTokenValid(token) {
// 	// 	fmt.Println("==================Access good")
// 	// 	next(w, r)
// 	// }
// 	refreshToken := service.RefreshTokens[refreshTokenFromRedis]

// 	//If refresh token's expiry is greater than 5 days but less than 7 days then new refresh token is generated
// 	if service.IsTokenValid(refreshToken.Exp, time.Hour*168) {
// 		newAccessToken := generateAccessUsingRefreshToken(refreshTokenFromRedis)

// 		fmt.Println("----------------Access generated -----------------")

// 		var tokens []byte
// 		if !service.IsTokenValid(refreshTokenFromRedis, time.Hour*120) {
// 			newRefreshToken, err := service.GenerateRefreshToken()

// 			fmt.Println("RefreshToken generated ***************************")

// 			if err != nil {
// 				err = errors.Wrap(err, "Refresh token not generated.")
// 				log.Println(err)

// 			}
// 			tokens = getToken(newAccessToken, newRefreshToken)

// 			fmt.Println(tokens, "*********************")

// 			w.Write(tokens)
// 		}
// 		// w.Write([]byte())
// 	} else {
// 		http.Error(w, "Authorization failed", http.StatusUnauthorized)
// 	}

// }

//GenerateAccessUsingRefreshToken is used when previous access token has expired. It calls IsRefreshTokenValid to check if new access token can be generated.
func GenerateAccessUsingRefreshToken(refreshToken string) string {
	userUUID := service.RefreshTokens[refreshToken].UserUUID
	accessToken, err := service.GenerateAccessToken(userUUID)
	if err != nil {
		err = errors.Wrap(err, "Access token not generated")
		log.Println(err)
	}
	return accessToken
}

//FileInsideServer for testing
func FileInsideServer(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	w.Write([]byte("This is the file I am trying to access"))
}

//HelloController for testing
func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello controller"))
}
