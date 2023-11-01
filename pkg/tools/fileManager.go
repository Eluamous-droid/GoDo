package tools

import (
	"encoding/json"
	"os"

	"github.com/eluamous-droid/godo/pkg/models"
)

var filePath string = os.ExpandEnv("$HOME/.godo/")
var fileName string = "todos.txt"

func AppendToTasks(oldItems models.TodoItems, item models.TodoItem) models.TodoItems {
	var newTodos models.TodoItems
	newTodos.TodoItems = append(oldItems.TodoItems, item)
	return newTodos
}

func WriteTasksToFile(fileContent models.TodoItems) {
	ensureFileExists(filePath)
	f, err := os.OpenFile(filePath+fileName, os.O_WRONLY, 0600)
	check(err)
	defer f.Close()
	input, _ := json.Marshal(fileContent)
	clearFile(filePath + fileName)
	f.Write(input)
	f.Sync()
}

func ReadTodosFromFile() models.TodoItems {
	var todos models.TodoItems
	ensureFileExists(filePath)
	f, err := os.ReadFile(filePath + fileName)
	check(err)
	if !isFileEmpty(filePath + fileName) {
		err = json.Unmarshal(f, &todos)
		check(err)
	}
	return todos
}

func ensureFileExists(fpath string) {
	if _, err := os.Stat(filePath + fileName); err == nil {
		return
	}
	os.MkdirAll(filePath, 0755)
	os.Create(fpath + fileName)
}

func isFileEmpty(fPath string) bool {
	f, err := os.Stat(fPath)
	check(err)
	return f.Size() == 0
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
