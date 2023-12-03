package main

import (
	"fmt"
	"go-mongo-crud/controllers"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	router := httprouter.New();

	userController := controllers.NewUserController(getSession())

	// Routes
	router.GET("/users/:id", userController.GetUser)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	// Start the server
	http.ListenAndServe("localhost:8080", router)
	
}




// Connect returns a MongoDB session
func getSession() *mgo.Session {
    session, err := mgo.Dial("mongodb://localhost:27017")
    if err != nil {
		fmt.Println("Error connecting to MongoDB");
       panic(err)
    }
    return session;
}
