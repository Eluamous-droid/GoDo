package tools

import (
	"flag"
	"os"
)


func ReadInput(){

    args := os.Args[1:]

    switch args[0] {
    case "list":
        ReadTodos()
    default:
    SaveToFile(args[0])
    }

    flag.Parse()
	
}