package service

import (
	"encoding/json"
	"fmt"

	"github.com/bhupeshbhatia/go-authserver/models"
	"github.com/gofrs/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Access token test", func() {
	Context("Generate access token", func() {
		var (
			generatedUUID uuid.UUID
			UUIDByteArray []byte
			uniqueID      string
			user          models.User
		)

		BeforeEach(func() {
			//Generate UUID
			generatedUUID, err := uuid.NewV4()
			Expect(err).ToNot(HaveOccurred())

			UUIDByteArray, err := generatedUUID.MarshalText()
			Expect(err).ToNot(HaveOccurred())

			uniqueID := string(UUIDByteArray)

			user := models.User{
				UUID:     uniqueID,
				Username: "test",
				Password: "testing",
			}
		})

		It("should return token string with UUID", func() {
			token, err := GenerateAccessToken(uniqueID)
			Expect(err).ToNot(HaveOccurred())

			t := &models.AccessToken{}
			err := json.Unmarshal(token, t)
			t.exp
			fmt.Println(token)

		})

	})
})
