package Manager

import (
	"flag"
	"fmt"
	"github.com/konojunya/cget/Utils"
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
		getCodeUrl(items[1])
	default:
		notfound()
	}
}

func showAllFile() {
	for _, key := range Utils.GetMenuKeys() {
		fmt.Printf("%s\n", key)
	}
}

func getCodeUrl(key string) {
	url := Utils.GetCodeUrl(key)
	code := Utils.FetchCode(url)

	fmt.Printf("Downloading %s from %s...\n",key,url)

	Utils.Export(key, code)
}

func help() {
	fmt.Println("command list.")
	fmt.Println("$ cget list\t\tYou can get code list.")
	fmt.Println("$ cget fetch [key]\tYou can save file from key.")
}

func notfound() {
	fmt.Println("this args is not found.")

	fmt.Println("$ cget list\t\tYou can get code list.")
	fmt.Println("$ cget fetch [key]\tYou can save file from key.")
	fmt.Println("$ cget help\t\thelp.")
}
