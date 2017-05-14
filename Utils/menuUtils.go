package Utils

import (
	"log"
	"encoding/json"
	"io/ioutil"
)

type Menu struct {
	Key string
	Url string
}

func initMenu() []Menu {

	bytes, err := ioutil.ReadFile("menu.json")
	if err != nil { log.Fatal(err) }

	var menus []Menu
	if err := json.Unmarshal(bytes, &menus); err != nil {
		log.Fatal(err)
	}

	return menus

}

func GetMenuKey() []string {
	menus := initMenu()

	var menuKeys []string

	for _, item := range menus{
		menuKeys = append(menuKeys,item.Key)
	}

	return menuKeys
}