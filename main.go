package main

import (
	"fmt"
	"github.com/konojunya/get-code-from-github/Utils"
)

func main() {
	code := Utils.FetchCode("https://raw.githubusercontent.com/Lebe-Inc/react-handson-vol2/blank/level2/webpack.config.js")
	fmt.Printf("%s", code)

	for _, item := range Utils.GetMenuKey(){
		fmt.Printf("%s\n",item)
	}
}
