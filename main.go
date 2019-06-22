package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kevinyang777/loginer/controller"
)

func main() {

	router := gin.Default()

	router.POST("/api/user", controller.CreateAccount)
	// router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	// router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	// router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	// router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler
	// router.run()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
