package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dhpsagala/crowdfunding_donny/controllers"
	"github.com/dhpsagala/crowdfunding_donny/libs"
	"github.com/dhpsagala/crowdfunding_donny/models"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.InitDb()
	router := mux.NewRouter()

	router.HandleFunc("/api", controllers.Index).Methods("GET")
	router.HandleFunc("/api/healthcheck", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/api/token", controllers.AuthenticateUser).Methods("POST")
	router.HandleFunc("/api/users/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/transaction", libs.AuthUser(controllers.UserTransaction)).Methods("GET")
	router.HandleFunc("/api/users/expense", libs.AuthUser(controllers.UserExpense)).Methods("GET")
	router.HandleFunc("/api/products", libs.AuthUser(controllers.ListOfAvailableItems)).Methods("GET")
	router.HandleFunc("/api/products/{id}/buy", libs.AuthUser(controllers.BuyProduct)).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println(port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Print(err)
	}
}
