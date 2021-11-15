package tools

import (
	"fmt"
	"os"
	"path"
)

var filePath string = os.ExpandEnv("$HOME/.todocli/")
var fileName string = "todos.txt"

func SaveToFile(input string){
	
	ensureBaseDir(filePath)
	f,err := os.OpenFile(filePath+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()

	_,err2 := f.WriteString(input + " \n")

	check(err2)
	f.Sync()
}

func ReadTodos(){
	ensureBaseDir(filePath)
	f,err := os.ReadFile(filePath+fileName)
	check(err)

	fmt.Print(string(f))

}

func ensureBaseDir(fpath string) error {
    baseDir := path.Dir(filePath)
    info, err := os.Stat(baseDir)
    if err == nil && info.IsDir() {
       return nil
    }
    return os.MkdirAll(baseDir, 0755)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}