package Utils

import (
	_"log"
	_"os"
	"fmt"
	"github.com/konojunya/cget/Constants"
)


func Export(exportData *Constants.ExportFormat) {
	//crr, _ := os.Getwd()
	//file, err := os.Create(crr + "/" + filename)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//
	//file.Write(code)

	fmt.Println(&exportData)
}
