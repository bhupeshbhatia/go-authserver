package service

// import (
// 	"log"
// 	"time"

// 	"github.com/bhupeshbhatia/go-authserver/models"
// 	jwt "github.com/dgrijalva/jwt-go"
// 	"github.com/gofrs/uuid"
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	"github.com/pkg/errors"
// )

// var _ = Describe("Redis client test", func() {
// 	Context("Redis db connection for storing refresh tokens", func() {
// 		var redis *service.Redis
// 		var client *redis.Client
// 		var refreshToken *models.RefreshToken
// 		var userUUID string

// 		BeforeEach(func() {
// 			type RedisClientTest struct{}
// 			var (
// 				add      string = "localhost:6379"
// 				password        = ""
// 				db              = 0
// 			)

// 			client, err := redis.RedisClient(add, password, db)
// 			Expect(err).ToNot(HaveOccurred())
// 		})

// 		It("Sets the data", func() {
// 			token := jwt.New(jwt.SigningMethodRS512)

// 			//This is where the token is signing with the private key in the backend
// 			tokenString, err := token.SignedString(getPrivateKey())
// 			Expect(err).ToNot(HaveOccurred())

// 			//Convert UUID from string
// 			generatedUUID, err := uuid.NewV4()
// 			if err != nil {
// 				err = errors.Wrap(err, "UUID not generated.")
// 				log.Println(err)
// 			}

// 			//Convert UUID to byte[]
// 			convertedByteArrayUUID, err := generatedUUID.MarshalText()
// 			if err != nil {
// 				err = errors.Wrap(err, "UUID not converted to byte[].")
// 				log.Println(err)
// 			}

// 			//Changed byte[] to string
// 			userUUID = string(convertedByteArrayUUID)

// 			refreshToken = models.RefreshToken{
// 				UserUUID: userUUID,
// 				Exp:      time.Now().AddDate(0, 0, 7).String(),
// 				Token:    tokenString,
// 			}

// 			setToken, err := redis.SetToken(userUUID, refreshToken, client)
// 			Expect(err).ToNot(HaveOccurred())
// 			Expect(setToken).To(Equal("Token added"))
// 		})

// 		It("Gets the data", func() {
// 			val, err := redis.GetToken(userUUID, client)
// 			Expect(err).ToNot(HaveOccured())
// 			Expect(val).To(Equal(refreshToken)) //Doesn't work
// 		})
// 	})
// })
