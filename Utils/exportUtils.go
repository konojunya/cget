package Utils

import (
	"log"
	"os"
	"github.com/konojunya/cget/Constants"
	"fmt"
)


func Export(exportData *Constants.ExportFormat) {
	crr, _ := os.Getwd()
	file, err := os.Create(crr + "/" + exportData.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(exportData.Code)

	fmt.Printf("Export %s\n",exportData.Filename)
}
