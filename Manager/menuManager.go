package Manager

import (
	"flag"
	"fmt"
	"github.com/konojunya/cget/Utils"
	"github.com/konojunya/cget/Constants"
)

func Init() {

	flag.Parse()
	items := flag.Args()

	switch items[0] {
	case "help":
		help()
	case "list":
		showAllFile()
	case "fetch":
		fetchCodes(items[1:])
	default:
		not_found()
	}
}

func showAllFile() {
	for _, key := range Utils.GetMenuKeys() {
		fmt.Printf("%s\n", key)
	}
}

func fetchCodes(keys []string){

	statusChan := getCodeFromGithub(keys)

	for range keys{
		select {
		case status := <-statusChan:
			Utils.Export(status)
		}
	}
}

func getCodeFromGithub(keys []string) <-chan *Constants.ExportFormat {

	statusChan := make(chan *Constants.ExportFormat)

	for _,key := range keys{
		go func(key string){
			url := Utils.GetCodeUrl(key)
			code := Utils.FetchCode(url)

			statusChan <- &Constants.ExportFormat{
				Filename: key,
				Code: code,
			}
		}(key)
	}

	return statusChan

}

func help() {
	fmt.Println("command list.")
	fmt.Println("$ cget list\t\tYou can get code list.")
	fmt.Println("$ cget fetch [key]\tYou can save file from key.")
}

func not_found() {
	fmt.Println("this args is not found.")

	fmt.Println("$ cget list\t\tYou can get code list.")
	fmt.Println("$ cget fetch [key]\tYou can save file from key.")
	fmt.Println("$ cget help\t\thelp.")
}
