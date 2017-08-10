package cache

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ieee0824/getenv"
)

var CACHE_DIR = fmt.Sprintf("%v/.tanya", getenv.String("HOME"))

var cache = map[string][]byte{}

func init() {
	if err := os.Mkdir(CACHE_DIR, 0700); err != nil {
		if err.Error() != fmt.Sprintf("mkdir %v: file exists", CACHE_DIR) {
			log.Fatalln(err)
		}
	}
}

func GetKey() []byte {
	bin, err := ioutil.ReadFile(fmt.Sprintf("%v/key", CACHE_DIR))
	if err != nil {
		if err.Error() != fmt.Sprintf("open %v/key: no such file or directory", CACHE_DIR) {
			log.Fatalln(err)
		}
		key := make([]byte, 32)
		if _, err := rand.Read(key); err != nil {
			log.Fatalln(err)
		}
		if err := ioutil.WriteFile(fmt.Sprintf("%v/key", CACHE_DIR), key, 0600); err != nil {
			log.Fatalln(errors.New("can not save key"))
		}

		return key
	}
	return bin
}

func LoadCache() {
	bin, err := ioutil.ReadFile(fmt.Sprintf("%v/cache", CACHE_DIR))
	if err != nil {
		if err.Error() != fmt.Sprintf("open %v/cache: no such file or directory", CACHE_DIR) {
			log.Fatalln(err)
		}
		return
	}

	if err := json.Unmarshal(bin, &cache); err != nil {
		log.Fatalln(err)
	}
}
