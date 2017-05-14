package Utils

import (
	"log"
	"os"
)

func Export(filename string, code []byte) {
	crr, _ := os.Getwd()
	file, err := os.Create(crr + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(code)
}
