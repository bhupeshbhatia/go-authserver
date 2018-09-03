package service

import (
	"log"

	"github.com/bhupeshbhatia/go-authserver/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("Jwt", func() {
	Context("Jwt access and refresh tokens are created", func() {

		// type user struct {
		// 	UUID     string `json:"uuid"`
		// 	Username string `json:"username"`
		// 	Password string `json:"password"`
		// }

		//generating UUID
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

		newUser := models.User{
			UUID:     uniqueID,
			Username: "test",
			Password: "testing",
		}

		It("should return error if UUID is empty", func() {
			_, err := GenerateAccessToken("")
			Expect(err).To(HaveOccurred())
		})

		It("should return error if UUID is incorrect", func() {
			_, err := GenerateAccessToken("testing")
			Expect(err).To(HaveOccurred())
		})

		It("should return token string with UUID", func() {
			token, err := GenerateAccessToken(uniqueID)
			Expect(err).ToNot(HaveOccurred())
			Expect(token).To(HaveOccurred())

		})

		It("should return empty user model", func() {
			testUser := models.User{
				UUID:     "",
				Username: "",
				Password: "",
			}
			userModel, err := AuthenticateUser(&testUser)
			Expect(err).To(HaveOccurred())
			Expect(userModel).Should(BeNil())
		})

		It("should return user model", func() {
			testUser := models.User{
				UUID:     uniqueID,
				Username: "test",
				Password: "testing",
			}
			userModel, err := AuthenticateUser(&testUser)
			Expect(err).ToNot(HaveOccurred())
			Expect(userModel).To(Should(BeEquivalentTo(&testUser)))
		})

		It("should return a refresh token", func() {
			token := jwt.New(jwt.SigningMethodRS512)

			//This is where the token is signing with the private key in the backend
			tokenString, err := token.SignedString(getPrivateKey())
			if err != nil {
				err = errors.Wrap(err, "Refresh token not signed.")
				log.Println(err)
				return "", err
			}
			refreshToken, err := GenerateRefreshToken()
			Expect(err).ToNot(HaveOccurred())
			Expect(refreshToken).Should(BeEquivalentTo(tokenString))
		})

	})
})
