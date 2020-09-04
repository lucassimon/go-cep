package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/tayron/go-cep/application/service"
	"github.com/stretchr/testify/assert"
)

func removeCacheFile(t *testing.T, id string) {
	assert.Nil(t, os.Remove(service.GetCacheFilename(id)))
}

func Test_getCacheFilename(t *testing.T) {
	id := "89201405"
	idWithDash := "89201-405"

	assert.Equal(t, service.GetCacheFilename(id), os.TempDir()+"/cep"+id)
	assert.Equal(t, service.GetCacheFilename(idWithDash), os.TempDir()+"/cep"+id)
}

func Test_getCep(t *testing.T) {
	t.Run("invalid cep", func(t *testing.T) {
		id := "0000000"
		wrongCep, err := service.GetCep(id)
		assert.Equal(t, "", wrongCep)
		assert.Error(t, err)
	})
	t.Run("valid cep", func(t *testing.T) {

		id := "60170150"
		cepJSON, err := service.GetCep(id)
		assert.Nil(t, err)
		res := service.Cep{}
		assert.Nil(t, json.Unmarshal([]byte(cepJSON), &res))

		assert.Equal(t, "60170-150", res.Cep)
		assert.Equal(t, "Rua Vicente Leite", res.Logradouro)
		assert.Equal(t, "até 879/880", res.Complemento)
		assert.Equal(t, "Meireles", res.Bairro)
		assert.Equal(t, "Fortaleza", res.Localidade)
		assert.Equal(t, "CE", res.Uf)
		assert.Equal(t, "", res.Unidade)
		assert.Equal(t, "2304400", res.Ibge)
		assert.Equal(t, "", res.Gia)

		removeCacheFile(t, id)
	})
}

func Test_Cache(t *testing.T) {
	id := "89201405"
	_, err := service.GetCep(id) // Add to temporary_directory_path/cep89201405
	assert.Nil(t, err)
	_, err = os.Stat(service.GetCacheFilename(id))
	assert.Nil(t, err)
	removeCacheFile(t, id)
}
