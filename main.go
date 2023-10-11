package main

import (
	"net/http"

	"github.com/FelipeGeraldoblufus/goapi/db"
	"github.com/FelipeGeraldoblufus/goapi/models"
	"github.com/FelipeGeraldoblufus/goapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Product{})
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Mail{})

	r := mux.NewRouter()

	r.HandleFunc("/api/mail", routes.PostMail).Methods("POST")
	r.HandleFunc("/api/mail/{id}", routes.GetMail).Methods("GET")
	r.HandleFunc("/api/mail/{id}", routes.UpdateMail).Methods("PUT")
	r.HandleFunc("/api/mail/{id}", routes.DeleteMail).Methods("DELETE")

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUser).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Products routes
	r.HandleFunc("/products", routes.GetProductsHandler).Methods("GET")
	r.HandleFunc("/products", routes.CreateProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", routes.GetProductHandler).Methods("GET")
	r.HandleFunc("/products/{id}", routes.DeleteProductHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

	//SendEmail.EmailConnect()

}
