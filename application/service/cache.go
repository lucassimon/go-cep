package service

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/ryanuber/go-filecache"
)

const cacheTime = 500

func GetFromCache(id string) string {
	updater := func(path string) error {
		return errors.New("expired")
	}

	fc := filecache.New(GetCacheFilename(id), cacheTime*time.Second, updater)

	fh, err := fc.Get()
	if err != nil {
		return ""
	}

	content, err := ioutil.ReadAll(fh)
	if err != nil {
		return ""
	}

	return string(content)
}

func SaveOnCache(id string, content string) string {
	updater := func(path string) error {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write([]byte(content))
		return err
	}

	fc := filecache.New(GetCacheFilename(id), cacheTime*time.Second, updater)

	_, err := fc.Get()
	if err != nil {
		return ""
	}

	return content
}

func GetCacheFilename(id string) string {
	return os.TempDir() + "/cep" + strings.Replace(id, "-", "", -1)
}
