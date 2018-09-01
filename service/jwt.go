package service

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhupeshbhatia/go-authserver/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//JWTAuthentication files struct
type JWTAuthentication struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// //InitJWTAuthentication will return backend instance
// func InitJWTAuthentication() *JWTAuthentication {
// 	authInstance := &JWTAuthentication{
// 		privateKey: getPrivateKey(),
// 		PublicKey:  getPublicKey(),
// 	}
// 	return authInstance
// }

func getPrivateKey() *rsa.PrivateKey {

	pKey := `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4w5xhil8YFSLptRxzQsiJgQm7DxfVx7nEFAndQDw/7a1VfIf
hhzZlUYx6u+57kP4+JPhqLMl9hEPnJh2DMPV4wrQAOSe6pDK5UP/xZQx8ygy70lG
fJ6MVo7mkXKaofKobOhkFIOhqtLU/6CrzFl+KdFIsD7pt+FxV6mMmPbnAvDN+hF5
NwU6N61WGAZER8z7SSTgayGpuHdUKCdPwfuiUIEX3GxhskzV/ROiS+R/NbQZlsfm
QqcBJ5FxhOtAVevi9s7x6LLTSQKopuuunSTTtu3ys/hs5m6AqNPPkLKqp6R8iXF1
Lg0DMeQlFHYwEo3oRweMNhfYRzC3ukioSf+GuwIDAQABAoIBADlemeKLMujoE80Y
WpSzXnJ6lBcWfgR2Q23EwuN2VG5YDONlZP+u5G8qKEyzO6hvNkYgn2DPuyS8VNR9
VT6OcMmIHtxK57he01UwZDzY3/IPUydQvWWZbd4lBy7y5Q1MUbAK29avF7cgxD6+
qwncBtusDJCzpLwYU1oR9ftkTyRXl8WzHUQ+/QILNnSCDsTrP8JsVaVxbd6FhKKn
5sSyqM+dX7mtvVAOcj0OJSHZiit7fk5QG9Pi/5iP4pCdZf42sImsr++2GFOezfJd
H5UU+ujTf+b4oGirnqgEDRrSr5IyykagWc07D2KJgyPzrkfFDxoB5C/ZC3C6C9AA
Xwzd+GECgYEA5SPDfCMVBRFkYBoxKgbWEElquGiPMDSe+p6QSlX24UXFv8gzdtbT
f33d27v2cpIOWYym3Er5JiSFq6oCr1cg9+mLP/tNc50sHrdHb8vRfn190nawFJHa
eOe0b3ZePUtAxdd1HaZgq4bNnLYSbi//spdHuu6E1jZrzcmbvIm7PJECgYEA/awp
rILMDvqHuGNlVr+kdcGfmFxA8y9Z1tZHLgqNjPQQlaOuyJn1cfYbIqghMLjk//Au
VQ5gfKLc2abHQaVQ2dLqV846eNQvr+cnLQUrUqk41IZuN0HTMbvLHgOLkQNdsUMs
1TmmPeMxh9X9cLqp7mZoY5CeWeWFOe3EJA1dZIsCgYEAklbf3yUMpJrx7wprQbrx
9Z7dwH5OjGve6JJh9oemT0LfQ1dZvtj+ZBr/mPkXMR6keX6Bhol/S2Ph1ruSUWck
0A/gdfFKCr9jUQ6eWgDif5UnyUUxuUFZNQRN0S3Yi+7GpFOxIUmDzagfIqmJZcPT
2rwQ/IqeXayN9vR+ONABu3ECgYAECn4PdXXytyL6WPsASsU/6vmz36RZO2Pe/ELe
BOUEXc7100mxgGJckmMURkFhGVDsktLqH/SBh8ak4PdDoHKNRcLd6zcbPaYU00XY
fcCW7IMvP4T59F586FTwAXZztO4FKODJ9MUlLz1WwJ3s8cxLM+5tx5v+Kp3YsmTx
fhUCyQKBgDCEkFexrqC2a1rHLh+pwTyvnE4JCVNt72FF8L51aEsG5tGGFvTvgUN6
IlRCYASNhUK/3+hu337uOSolKXu0W+dFnp1/OLo6sUkuhxWGx3YLwGJygjSrOl5f
3wIikQ0U/RjRr+/pI0/yw/w3Xcr7iUjei6SBxkiIeZL/749EcLNB
-----END RSA PRIVATE KEY-----`

	data, _ := pem.Decode([]byte(pKey))

	// privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {

		err = errors.Wrap(err, "")
		// log.Println(utils.ErrorStackTrace(err))
		log.Fatalln(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {

	pKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4w5xhil8YFSLptRxzQsi
JgQm7DxfVx7nEFAndQDw/7a1VfIfhhzZlUYx6u+57kP4+JPhqLMl9hEPnJh2DMPV
4wrQAOSe6pDK5UP/xZQx8ygy70lGfJ6MVo7mkXKaofKobOhkFIOhqtLU/6CrzFl+
KdFIsD7pt+FxV6mMmPbnAvDN+hF5NwU6N61WGAZER8z7SSTgayGpuHdUKCdPwfui
UIEX3GxhskzV/ROiS+R/NbQZlsfmQqcBJ5FxhOtAVevi9s7x6LLTSQKopuuunSTT
tu3ys/hs5m6AqNPPkLKqp6R8iXF1Lg0DMeQlFHYwEo3oRweMNhfYRzC3ukioSf+G
uwIDAQAB
-----END PUBLIC KEY-----`

	data, _ := pem.Decode([]byte(pKey))

	// publicKeyFile.Close()

	publicKeyImported, error := x509.ParsePKIXPublicKey(data.Bytes)
	if error != nil {
		log.Println(error)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
	if !ok {
		log.Println(error)
	}

	return rsaPub
}

//GenerateAccessToken creates the token for JWT authentication
func GenerateAccessToken(userUUID string) (string, error) {
	//When checking the token - this is the algorithm used
	token := jwt.New(jwt.SigningMethodRS512)

	//Claims on jwt tokens
	/*
			Audience  string `json:"aud,omitempty"`
		    ExpiresAt int64  `json:"exp,omitempty"`
		    Id        string `json:"jti,omitempty"`
		    IssuedAt  int64  `json:"iat,omitempty"`
		    Issuer    string `json:"iss,omitempty"`
		    NotBefore int64  `json:"nbf,omitempty"`
		    Subject   string `json:"sub,omitempty"`

	*/
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 15).Unix(),
		"iat": time.Now().Unix(),
		"jti": userUUID,
	}

	//This is where the token is signing with the private key in the backend
	tokenString, err := token.SignedString(getPrivateKey())
	if err != nil {
		err = errors.Wrap(err, "JWT is not signed.")
		return "", err
	}

	return tokenString, nil
}

//AuthenticateUser connects to MongoDb here. Need to come up with a better way of passing passwords
func AuthenticateUser(user *models.User) bool {
	//This is only temperary -- Should technically check from Db
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)

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

	testUser := models.User{
		UUID:     uniqueID,
		Username: "test",
		Password: string(hashedPassword),
	}
	log.Println(user)
	log.Println(testUser)
	log.Println(bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)))

	fmt.Println(testUser.UUID, testUser.Username, testUser.Password)

	return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
}

//RefreshToken creates the refresh token for JWT authentication
func RefreshToken() (string, error) {
	//When checking the token - this is the algorithm used
	token := jwt.New(jwt.SigningMethodRS512)

	//This is where the token is signing with the private key in the backend
	tokenString, err := token.SignedString(getPrivateKey())
	if err != nil {
		err = errors.Wrap(err, "Refresh token not signed.")
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}

//ParseAndDecryptToken service parses a request header and decrypts token
func ParseAndDecryptToken(r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return getPublicKey(), nil
	})

	if err != nil {
		err = errors.Wrap(err, "Error decrypting JWT")
		log.Println(err)
	}
	return token, nil
}

//IsAccessTokenValid checks the expiry time of the token
func IsAccessTokenValid(token *jwt.Token) bool {
	var expiryTime int64
	//Check the expiry time in claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expiryTime = claims["exp"].(int64)
	}

	if expiryTime < time.Now().Unix() {
		return true
	}
	return false
}
