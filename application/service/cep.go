package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

func GetCep(id string) (string, error) {
	cached := GetFromCache(id)
	if cached != "" {
		return cached, nil
	}
	req, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", id))
	if err != nil {
		return "", err
	}

	var c Cep
	err = json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		return "", err
	}
	res, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return SaveOnCache(id, string(res)), nil
}
