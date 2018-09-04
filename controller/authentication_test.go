package controller

import (
	"net/http"

	"github.com/bhupeshbhatia/go-authserver/controller"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/urfave/negroni"
)

var _ = Describe("Endpoint access", func() {
	Context("Login, Generate tokens and logout endpoints", func() {

		var (
			endPoint *EndPointAccess
			token    []byte
			router   *mux.Router
		)

		BeforeEach(func() {
			router := mux.NewRouter()
			n := negroni.Classic()
			n.UseHandler(router)
			http.ListenAndServe(":8080", n)

		})

		It("should return status unauthorized if user information is absent ", func() {
			router.HandleFunc("/login", controller.Login).Methods("POST")

			Expect(err).To(HaveOccurred())
		})

	})
})
