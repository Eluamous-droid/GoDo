package tools

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var filePath string = os.ExpandEnv("$HOME/.todocli/")
var fileName string = "todos.txt"

func AppendToFile(input string){
	
	ensureBaseDir(filePath)
	f,err := os.OpenFile(filePath+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()

	_,err2 := f.WriteString(input + " \n")

	check(err2)
	f.Sync()
}

func writeToFile(input string){
	ensureBaseDir(filePath)

	clearFile(filePath+fileName)
	f,err := os.OpenFile(filePath+fileName, os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()

	_,err2 := f.WriteString(input + " \n")

	check(err2)
	f.Sync()
}


func readTodosFromFile() []string {
	ensureBaseDir(filePath)
	f,err := os.ReadFile(filePath+fileName)
	check(err)
	todos := strings.Split(string(f),"\n")
	todos = todos[0:len(todos)-1]
	return todos
}

func PrintTodos(){
	todos :=readTodosFromFile()
	for i := 0; i < len(todos) - 1; i++{
		fmt.Println(i," "+todos[i])
	}
}

func DeleteFromFile(index int){
	todos := readTodosFromFile()
	todos = removeIndex(todos, index)
	writeToFile(buildString(todos))
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func buildString(input []string) string{
	output := ""
	for _,s := range input{
		output += s + "\n"
	}
	return output
}

func ensureBaseDir(fpath string) error {
    baseDir := path.Dir(filePath)
    info, err := os.Stat(baseDir)
    if err == nil && info.IsDir() {
       return nil
    }
    return os.MkdirAll(baseDir, 0755)
}

func clearFile(fpath string){
	err := (os.Truncate(fpath,0))
	check(err)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}