package Manager

import (
	"flag"
	"fmt"
	"github.com/konojunya/get-code-from-github/Utils"
)

func Init() {

	flag.Parse()
	items := flag.Args()

	switch items[0] {
	case "help":
		help()
	case "list":
		showAllMenu()
	case "get":
		getCodeUrl(items[1])
	default:
		notfound()
	}
}

func showAllMenu() {
	for _, key := range Utils.GetMenuKeys() {
		fmt.Printf("%s\n", key)
	}
}

func getCodeUrl(key string) {
	url := Utils.GetCodeUrl(key)
	code := Utils.FetchCode(url)

	Utils.Export(key,code)
}

func help() {}

func notfound() {
	fmt.Println("this args is not found.")
}
