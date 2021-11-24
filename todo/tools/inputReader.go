package tools

import (
	"flag"
	"os"
	"strconv"
)


func ReadInput(){

    args := os.Args[1:]

    switch args[0] {
    case "list":
        PrintTodos()
    case "delete":
        if input, err := strconv.Atoi(args[1]); err == nil{
        DeleteFromFile(input)
        } else {
            println("Please put the index of the todo youre trying to delete")
        }
    default:
    AppendToFile(args[0])
    }

    flag.Parse()
	
}
