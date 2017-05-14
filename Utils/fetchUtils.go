package Utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

func FetchCode(url string) []byte {

	fmt.Printf("Downloading %s from github.com...\n",url)

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
