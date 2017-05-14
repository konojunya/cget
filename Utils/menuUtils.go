package Utils

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
)

type Menu struct {
	Key string
	Url string
}

func initMenu() []Menu {

	res, err := http.Get("https://raw.githubusercontent.com/konojunya/cget/master/menu.json")
	if err != nil { log.Fatal(err) }

	bytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil { log.Fatal(err) }

	var menus []Menu
	if err := json.Unmarshal(bytes, &menus); err != nil {
		log.Fatal(err)
	}

	return menus

}

func GetMenuKeys() []string {
	menus := initMenu()

	var menuKeys []string

	for _, item := range menus {
		menuKeys = append(menuKeys, item.Key)
	}

	return menuKeys
}

func GetCodeUrl(key string) string {

	menus := initMenu()

	var url string

	for _, item := range menus {
		if item.Key == key {
			url = item.Url
		}
	}

	return url
}
