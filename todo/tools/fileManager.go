package tools

import (
	"os"
	"path"
)

func SaveToFile(input string){
	path := "$HOME/.todocli/"
	ensureBaseDir(path)
	f,err := os.Create(os.ExpandEnv(path+"todos.txt"))
	check(err)
	defer f.Close()

	_,err2 := f.WriteString(input)

	check(err2)
	f.Sync()


}

func ensureBaseDir(fpath string) error {
    baseDir := path.Dir(fpath)
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