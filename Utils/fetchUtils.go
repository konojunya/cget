package Utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

func FetchCode(url string) []byte {

	if url == "" {
		fmt.Println("url is not selected!")
	}

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
