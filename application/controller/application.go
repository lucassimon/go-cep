package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tayron/go-cep/application/service"
)

const errorMessage = "Erro lendo CEP"

func Home(rw http.ResponseWriter, req *http.Request) {
	_, err := rw.Write([]byte([]byte("ping")))
	if err != nil {
		respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
		return
	}
}

func GetCep(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	rw.Header().Set("Content-Type", "application/json")

	cepCleaned := strings.ReplaceAll(vars["id"], "-", "")
	cepCleaned = strings.ReplaceAll(cepCleaned, ".", "")

	cep, err := service.GetCep(cepCleaned)
	if err != nil {
		respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
		return
	}
	_, err = rw.Write([]byte(cep))
	if err != nil {
		respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
		return
	}
}

//RespondWithError return a http error
func respondWithError(w http.ResponseWriter, code int, e string, message string) {
	respondWithJSON(w, code, map[string]string{"code": strconv.Itoa(code), "error": e, "message": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
