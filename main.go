package main

import (
	"fmt"
	"github.com/dhpsagala/crowdfunding_donny/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api", controllers.Index).Methods("GET")
	router.HandleFunc("/api/healthcheck", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/api/token", controllers.AuthenticateUser).Methods("POST")
	router.HandleFunc("/api/user/register", controllers.CreateUser).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println(port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Print(err)
	}
}
