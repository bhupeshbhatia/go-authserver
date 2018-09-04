package service

import (
	"fmt"
	"log"
	"time"

	"github.com/bhupeshbhatia/go-authserver/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// var RefreshTokens map[string]string
var RefreshTokens = make(map[string]models.RefreshToken)

type Authentication struct {
	Time    time.Time
	NowTime func() time.Time
}

func (a *Authentication) GenerateAccessToken(userUUID string) (string, error) {
	if userUUID == "" {
		return "", errors.New("User-UUID not set")
	}
	//When checking the token - this is the algorithm used
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 15).Unix(),
		"iat": time.Now().Unix(),
		"jti": userUUID,
	}

	//This is where the token is signing with the private key in the backend
	tokenString, err := token.SignedString(a.getPrivateKey())

	if err != nil {
		err = errors.Wrap(err, "JWT is not signed.")
		return "", err
	}

	return tokenString, nil
}

//AuthenticateUser connects to MongoDb here. Need to come up with a better way of passing passwords
func (a *Authentication) AuthenticateUser(user *models.User) (*models.User, error) {
	// calls db
	// 1. UUID - user

	//This is only temperary -- Should technically check from Db
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

	fmt.Println(uniqueID)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)

	testUser := models.User{
		UUID:     uniqueID,
		Username: "test",
		Password: string(hashedPassword),
	}
	log.Println(user)
	log.Println(testUser)
	log.Println(bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)))

	fmt.Println(testUser.UUID, testUser.Username, testUser.Password)

	isAuthenticated := user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil

	//See of we get the value of "token" in AccessToken method?
	if isAuthenticated {
		return &testUser, nil
	}
	return nil, errors.WithMessage(err, "User not found") //NEED TO RETURN ERROR
}

func (a *Authentication) createRefreshToken(userUUID string, token string) models.RefreshToken {
	refreshToken := models.RefreshToken{
		UserUUID: userUUID,
		Exp:      a.NowTime().AddDate(0, 0, 7),
		Token:    token,
	}
	return refreshToken
}

//GenerateRefreshToken creates the refresh token for JWT authentication
func (a *Authentication) GenerateRefreshToken(userUUID string) (string, error) {
	//When checking the token - this is the algorithm used
	token := jwt.New(jwt.SigningMethodRS512)

	//This is where the token is signing with the private key in the backend
	tokenString, err := token.SignedString(a.getPrivateKey())
	if err != nil {
		err = errors.Wrap(err, "Refresh token not signed.")
		log.Println(err)
		return "", err
	}

	//STORE IT IN DB with time = 7 days

	// refreshToken := models.RefreshToken{
	// 	UserUUID: userUUID,
	// 	Exp:      a.NowTime().AddDate(0, 0, 7),
	// 	Token:    tokenString,
	// }

	refreshToken := a.createRefreshToken(userUUID, tokenString)
	log.Println(refreshToken)

	// client := redis.NewClient(&redis.Options{
	// 	//"localhost:6379"
	// 	//""
	// 	//0
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// if err != nil {
	// 	return "", err
	// }

	// teest := client.Set(userUUID, &refreshToken, 6*time.Hour)
	// fmt.Println(teest)
	fmt.Println(tokenString)
	return tokenString, nil
}

//IsTokenValid checks if token the token has expired
func (a *Authentication) IsTokenValid(expiryTime time.Time, addDuration time.Duration) bool {
	//
	isExpired := expiryTime.Unix() <= time.Now().Add(addDuration).Unix()
	return isExpired

}
