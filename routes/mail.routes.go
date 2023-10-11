package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FelipeGeraldoblufus/goapi/SendEmail"
	"github.com/FelipeGeraldoblufus/goapi/db"
	"github.com/FelipeGeraldoblufus/goapi/models"
	"github.com/gorilla/mux"
)

func PostMail(w http.ResponseWriter, r *http.Request) {
	var mail models.Mail

	json.NewDecoder(r.Body).Decode(&mail)

	createmail := db.DB.Create(&mail) //se guarda para ver si existe
	err := createmail.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

	}
	// Establece el nombre y la dirección de correo electrónico del destinatario
	toName := mail.Rut // Cambia por el nombre adecuado
	toAddress := mail.Correo

	// Llama a la función EmailConnect para enviar el correo
	subjecto := "Comprobante de Pago"
	SendEmail.EmailConnect(toName, toAddress, subjecto)

	json.NewEncoder(w).Encode(&mail)
}

func GetMail(w http.ResponseWriter, r *http.Request) {
	var mail models.Mail
	params := mux.Vars(r)
	db.DB.First(&mail, params["id"])
	json.NewEncoder(w).Encode(&mail)

}

func UpdateMail(w http.ResponseWriter, r *http.Request) {
	// Manejar la actualización de un correo por su ID
	// Por ejemplo, puedes buscar el correo en la base de datos, actualizar sus propiedades y devolverlo como respuesta
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Actualizar correo con ID: %s", id)
}

func DeleteMail(w http.ResponseWriter, r *http.Request) {
	// Manejar la eliminación de un correo por su ID
	// Por ejemplo, puedes buscar el correo en la base de datos y eliminarlo, luego devolver una respuesta adecuada
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Eliminar correo con ID: %s", id)
}
