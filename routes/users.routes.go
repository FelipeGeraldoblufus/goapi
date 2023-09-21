package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelipeGeraldoblufus/goapi/db"
	"github.com/FelipeGeraldoblufus/goapi/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	db.DB.Model(&user).Association("Product").Find(&user.Product)

	json.NewEncoder(w).Encode(&user)

}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createUser := db.DB.Create(&user) //se guarda para ver si existe
	err := createUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)

}
