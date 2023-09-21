package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelipeGeraldoblufus/goapi/db"
	"github.com/FelipeGeraldoblufus/goapi/models"
	"github.com/gorilla/mux"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	db.DB.Find(&products)

	json.NewEncoder(w).Encode(products)

}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)
	db.DB.First(&product, params["id"])
	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("task not found"))
		return

	}
	json.NewEncoder(w).Encode(&product)

}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	createdproduct := db.DB.Create(&product)
	err := createdproduct.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}

	json.NewEncoder(w).Encode(&product)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)
	db.DB.First(&product, params["id"])
	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("task not found"))
		return

	}
	db.DB.Unscoped().Delete(&product)
	w.WriteHeader(http.StatusNoContent)
}
