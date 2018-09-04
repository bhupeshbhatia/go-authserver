package service

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/bhupeshbhatia/go-authserver/models"
// 	jwt "github.com/dgrijalva/jwt-go"
// 	"github.com/gofrs/uuid"
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// var TokensTests interface {
// 	GetTime(duration time.Duration) *time.Time
// 	CreateToken(userUUID string) string
// 	ParseToken(token string) *jwt.Token
// }

// var TokenAccess struct {
// 	Time *TokensTests
// }

// func (t *TokenAccess) GetTime(duration *time.Duration) *time.Time {
// 	return time.Now().Add(duration).Unix()
// }

// token := jwt.New(jwt.SigningMethodRS512)

// func (t *TokenAccess) CreateToken(userUUID string) string {
// 	//When checking the token - this is the algorithm used

// 	token.Claims = jwt.MapClaims{
// 		"exp": time.Now().Add(time.Minute * 15).Unix(),
// 		"iat": time.Now().Unix(),
// 		"jti": userUUID,
// 	}

// 	tokenString, err := token.SignedString(getPrivateKey())
// 	if err != nil {
// 		err = errors.Wrap(err, "JWT is not signed.")
// 		return ""
// 	}

// 	return tokenString
// }

// func (t *TokenAccess) ParseToken(tokenString string) *jwt.Token {
// 	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
// 		return getPublicKey()
// 	})
// }

// func ()

// var _ = Describe("Access token test", func() {
// 	Context("Generate access token", func() {
// 		var (
// 			generatedUUID uuid.UUID
// 			UUIDByteArray []byte
// 			uniqueID      string
// 			user          models.User
// 			var tTime *TokenAccess
// 		)

// 		BeforeEach(func() {
// 			//Generate UUID
// 			generatedUUID, err := uuid.NewV4()
// 			Expect(err).ToNot(HaveOccurred())

// 			UUIDByteArray, err := generatedUUID.MarshalText()
// 			Expect(err).ToNot(HaveOccurred())

// 			uniqueID := string(UUIDByteArray)

// 			user := models.User{
// 				UUID:     uniqueID,
// 				Username: "test",
// 				Password: "testing",
// 			}
// 		})

// 		It("should return token string with UUID", func() {
// 			token, err := GenerateAccessToken(uniqueID)
// 			Expect(err).ToNot(HaveOccurred())

// 			t := &models.AccessToken{}
// 			err := json.Unmarshal(token, t)

// 			fmt.Println(token)

// 		})
// 		It("checking time for token expiry", func(){
// 			timeNow :=tTime.GetTime(0)
// 			timeExp := tTime.GetTime(15)

// 			Expect(timeNow).shouldNot(Equal(timeExp))
// 		})

// 	})
// })
