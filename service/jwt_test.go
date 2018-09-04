package service

import (
	"testing"
	"time"

	"github.com/bhupeshbhatia/go-authserver/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Tests will use this testDatabase
func TestJWT(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JWT Suite")
}

var _ = Describe("Authentication service", func() {
	Context("Jwt access and refresh tokens are created", func() {

		// type user struct {
		// 	UUID     string `json:"uuid"`
		// 	Username string `json:"username"`
		// 	Password string `json:"password"`
		// }
		// var (
		// 	generatedUUID uuid.UUID
		// 	UUIDByteArray []byte
		// 	uniqueID      string
		// 	user          models.User
		// )

		// BeforeEach(func() {
		// 	//Generate UUID
		// 	generatedUUID, err := uuid.NewV4()
		// 	Expect(err).ToNot(HaveOccurred())

		// 	UUIDByteArray, err := generatedUUID.MarshalText()
		// 	Expect(err).ToNot(HaveOccurred())

		// 	uniqueID := string(UUIDByteArray)

		// 	user := models.User{
		// 		UUID:     uniqueID,
		// 		Username: "test",
		// 		Password: "testing",
		// 	}
		// })

		//================================

		It("blabla", func() {
			t := time.Now()
			mockTime := mocks.Time{
				Time: t,
			}

			nowTime := func() time.Time {
				return t
			}

			auth := Authentication{
				Time:    mockTime.Time,
				NowTime: nowTime,
			}

			token := auth.createRefreshToken("some-uuid", "asdfsdf")

			t = t.AddDate(0, 0, 7)
			Expect(token.Exp.Unix()).To(Equal(t.Unix()))
		})

		// ================================
		// It("should return error if UUID is empty", func() {
		// 	_, err := GenerateAccessToken("")
		// 	Expect(err).To(HaveOccurred())
		// })

		// It("should return error if UUID is incorrect", func() {
		// 	_, err := GenerateAccessToken("testing")
		// 	Expect(err).To(HaveOccurred())
		// })

		// It("should return empty user model", func() {
		// 	testUser := models.User{
		// 		UUID:     "",
		// 		Username: "",
		// 		Password: "",
		// 	}
		// 	userModel, err := AuthenticateUser(&testUser)
		// 	Expect(err).To(HaveOccurred())
		// 	Expect(userModel).Should(BeNil())
		// })

		// It("should return user model", func() {
		// 	// testUser := models.User{
		// 	// 	UUID:     uniqueID,
		// 	// 	Username: "test",
		// 	// 	Password: "testing",
		// 	// }
		// 	userModel, err := AuthenticateUser(&user)
		// 	Expect(err).ToNot(HaveOccurred())
		// 	// Expect(userModel).To(Should(BeEquivalentTo(&testUser)))
		// })

		// It("should return a refresh token", func() {
		// 	token := jwt.New(jwt.SigningMethodRS512)

		// 	//This is where the token is signing with the private key in the backend
		// 	tokenString, err := token.SignedString(getPrivateKey())
		// 	if err != nil {
		// 		err = errors.Wrap(err, "Refresh token not signed.")
		// 		log.Println(err)
		// 		return "", err
		// 	}
		// 	refreshToken, err := GenerateRefreshToken()
		// 	Expect(err).ToNot(HaveOccurred())
		// 	Expect(refreshToken).Should(BeEquivalentTo(tokenString))
		// })

	})
})
