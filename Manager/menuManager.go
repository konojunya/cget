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
	case "get":
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

	Utils.Export(key, code)
}

func help() {
	fmt.Println("command list.")
	fmt.Println("$ ./main list\t\tYou can get code list.")
	fmt.Println("$ ./main get [key]\tYou can save file from key.")
}

func notfound() {
	fmt.Println("this args is not found.")

	fmt.Println("$ ./main list\t\tYou can get code list.")
	fmt.Println("$ ./main get [key]\tYou can save file from key.")
	fmt.Println("$ ./main help\t\thelp.")
}
