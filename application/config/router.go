package config

import (
	"net/http"

	"github.com/tayron/go-cep/application/controller"

	"github.com/gorilla/mux"
)

func LoadHouter() {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Home)
	router.HandleFunc("/cep/{id}", controller.GetCep)
	http.Handle("/", router)
}
