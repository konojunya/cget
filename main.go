package main

import (
	_ "fmt"
	"github.com/konojunya/get-code-from-github/Manager"
	_ "github.com/konojunya/get-code-from-github/Utils"
)

func main() {
	//code := Utils.FetchCode("https://raw.githubusercontent.com/Lebe-Inc/react-handson-vol2/blank/level2/webpack.config.js")
	//fmt.Printf("%s", code)
	//
	//for _, item := range Utils.GetMenuKeys(){
	//	fmt.Printf("%s\n",item)
	//}

	//url := Utils.GetCodeUrl("webpack.config.js")
	//fmt.Printf(url)

	Manager.SelectAllMenu()

}
