package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/context"
	"github.com/joho/godotenv"
	"github.com/tayron/go-cep/application/config"
	"github.com/tayron/go-cep/application/util"
)

func init() {
	caminhoAplicacao := util.ObterCaminhoDiretorioAplicacao()

	environmentPath := filepath.Join(caminhoAplicacao, ".env")

	fmt.Println("Arquivo de configuração: " + environmentPath)
	err := godotenv.Load(environmentPath)

	if err != nil {
		panic(err)
	}
}

func main() {

	config.LoadHouter()

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + os.Getenv("LOCALHOST_PORT"),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	fmt.Println("Servidor executando no endereço: http://127.0.0.1:" + os.Getenv("LOCALHOST_PORT"))
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
