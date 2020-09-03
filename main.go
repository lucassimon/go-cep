package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eminetto/goCep/application/service"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

const serverPort = "3000"

func main() {
	errorMessage := "Erro lendo CEP"
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte([]byte("ping")))
		if err != nil {
			respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
			return
		}
	})
	router.HandleFunc("/cep/{id}", func(rw http.ResponseWriter, req *http.Request) {
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
	})
	http.Handle("/", router)
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + serverPort,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
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
