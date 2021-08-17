package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Buf []byte

func init() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Используем текущую директорию для поиска конфига")
		dir = "."
	}
	fn := filepath.Join(dir, "config.json")

	Buf, err = ioutil.ReadFile(fn)
	if err != nil {
		panic("Ошибка при чтении конфига config.json (отсутсвует файл?)")
	}
}

func Get(c interface{}) error {
	return json.Unmarshal(Buf, c)
}
