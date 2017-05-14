package Utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

func FetchCode(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	code, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return code

}
