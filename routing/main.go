package main

import (
	"net/http"

	"github.com/bhupeshbhatia/go-authserver/controller"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	router = setHelloRoute(router)
	router = setAuthenticationRoute(router)
	return router
}

func main() {
	router := initRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":8080", n)
}

func setAuthenticationRoute(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controller.Login).Methods("POST")
	// router.HandleFunc("/refresh-token", controller.).Methods("POST")

	router.Handle("/fileaccess",
		negroni.New(
			negroni.HandlerFunc(controller.ValidateAccessToken),
			negroni.HandlerFunc(controller.FileInsideServer),
		)).Methods("GET")

	// router.Handle("/logout",
	//     negroni.New(
	//         negroni.HandlerFunc(authentication.RequireTokenAuthentication),
	//         negroni.HandlerFunc(controllers.Logout),
	//     )).Methods("GET")

	return router
}

func setHelloRoute(router *mux.Router) *mux.Router {
	// router.Handle("/test/hello",
	// 	negroni.New(negroni.HandlerFunc(controller.ValidateAccessToken),
	// 		negroni.HandlerFunc(controller.HelloController),
	// 	)).Methods("GET")

	router.Handle("/hello",
		negroni.New(
			negroni.HandlerFunc(controller.ValidateTokens),
			negroni.HandlerFunc(controller.HelloController),
		)).Methods("GET")

	return router
}
