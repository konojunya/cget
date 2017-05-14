package Utils

import (
	"log"
	"os"
)

func Export(code []byte) {
	file, err := os.Create("webpack.config.js")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(code)
}
