package Utils

import (
	"log"
	"os"
)

func Export(filename string, code []byte) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(code)
}
