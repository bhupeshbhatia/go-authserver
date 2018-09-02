package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhupeshbhatia/go-authserver/models"
	"github.com/bhupeshbhatia/go-authserver/service"
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

	user := service.AuthenticateUser(requestUser)

	//UUID is empty --- cause there is nothing in the user --- I will generate a new one
	// Wait

	if user != nil {
		accessToken, err := service.GenerateAccessToken(user.UUID)
		if err != nil {
			err = errors.Wrap(err, "Access token not generated.")
			log.Println(err)
		}

		refreshToken, err := service.GenerateRefreshToken()
		if err != nil {
			err = errors.Wrap(err, "Refresh token not generated.")
			log.Println(err)
		}

		//
		exp, err := time.ParseDuration("168h")
		if err != nil {
			log.Fatalln(err)
		}
		service.RefreshTokens[refreshToken] = service.RefreshToken{
			UserUUID: user.UUID,
			Exp:      time.Now().Add(exp),
		}

		// Check the time value here?
		// I think it might be type converesion Somwhere in Json unmarsha;marshal
		//time: unknown unit d in duration 7d
		//error - crash
		// Try again? = login is good - iits because parseduration doesn't have d
		// Yup. Butwhat about /access-token? --- same thing

		// Everything works? - except for access token
		// Ah, lets figure it out tomorrow. I am having difficulty keeping my eyes open
		//Sure lets do it tomorrow//

		//Thanks for testing		//-62135596800 1535874416
		// That should do it?
		// Try again? -- crashed
		//..\controller\authentication.go:56:12: assignment mismatch: 2 variables but 1 values
		//..\controller\authentication.go:56:29: too many arguments in call to time.Now().Add

		tokens := getToken(accessToken, refreshToken)
		w.Header().Set("Content-Type", "application/json")
		w.Write(tokens)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func getToken(accessToken string, refreshToken string) []byte {
	tokens := Tokens{
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
	//What do you mean old? --- I thought we just checked the time?
	// Yeah, the refresh token that user gets at login, when it
	//expires, the client hits this endpoint to get a new one,
	// and passes the old one in return. Or did we discuss something else?
	// Lets not make refresh same as JWT, I might switch it with crypto-rand because of bandwidth efficiency. No need to encrypt/decrypt refresh token then?

	//what happens when token is empty?
	// 401
	// Yeah, we dont enrypt/add-secrets to refresh token. Just a random crypto string,
	// the way it is. Its pointless to add stuff to it,
	// since the key itself is the source of auth.
	//Do you want to add anything to refresh token look at jwt.go/ -> generateRefreshToken
	// I'll change that method to gnereate crypto/rand - a string.
	// Nothing else should be affected
	//What about after we receive refresh token? send it with accessToken?
	// on /refresh-token endpoint? Nah, just send refresh token.
	// I dont see why we need to send access token
	//Anything else for this?
	// Looks good to me. Gimme a min, I am just gonna create a new repo
	// for that events service so CI runs its tests in meantime....
	//OK
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	refreshToken, err := service.GenerateRefreshToken()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	w.Write([]byte(refreshToken))

}

//In login - access token is missing- its returning empty
// Any error? You got team viewer installed?

//AccessToken gets a new access token using RefreshToken
func AccessToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println("=====================")
	fmt.Println(token)
	//it crashed -- line 58
	// Error?
	// Wait, lemme try some stuff
	//..\controller\authentication.go:58:28: too many arguments in call to time.Now().Add
	// Oh, wait lol
	//FOrgot something
	refreshToken := service.RefreshTokens[token]

	fmt.Println(service.IsTokenValid(refreshToken.Exp, time.Hour*168)) //false
	// Wait.... Negative?!
	//refresh.exp.unix() = -62135596800
	// Are you sure that this oen is negaive ^ -- yup ran it again -- same num
	//time = 1535873877
	// Try again? <---
	fmt.Println(refreshToken.Exp.Unix(), time.Now().Unix()) //refresh is negative
	//time is higher
	// See which value is higher/lower?
	// This is the issue them.
	// How about this guy^
	if !service.IsTokenValid(refreshToken.Exp, time.Hour*168) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessToken, err := service.GenerateAccessToken(refreshToken.UserUUID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(accessToken))

}

//What's next? --- where do we check for 5 days thing?

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
