package tools

import (
	"encoding/json"
	"os"
	"path"

	"github.com/eluamous-droid/godo/pkg/models"
)

var filePath string = os.ExpandEnv("$HOME/.godo/")
var fileName string = "todos.txt"

func AppendToFile(input string) {

	ensureBaseDir(filePath)
	f, err := os.OpenFile(filePath+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()

	_, err2 := f.WriteString(input + " \n")
	check(err2)
	f.Sync()
}

func WriteToFile(input string) {
	ensureBaseDir(filePath)

	clearFile(filePath + fileName)
	f, err := os.OpenFile(filePath+fileName, os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()
	_, err2 := f.WriteString(input)

	check(err2)
	f.Sync()
}

func ReadTodosFromFile() []models.TodoItem {
	ensureBaseDir(filePath)
	f, err := os.ReadFile(filePath + fileName)
	check(err)

	var todos []models.TodoItem
	err = json.Unmarshal(f, &todos)
	check(err)
	//todos = todos[0:len(todos)-1]
	return todos
}

func ensureBaseDir(fpath string) error {
	baseDir := path.Dir(filePath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}

func clearFile(fpath string) {
	err := (os.Truncate(fpath, 0))
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
