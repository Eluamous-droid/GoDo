package tools

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)


func ReadInput(){

    args := os.Args[1:]

    switch args[0] {
    case "add": 
        if args[1] != ""{
            AppendToFile(args[1])
        }else{
            fmt.Println("Missing todo to add")
        }
    case "delete":
        if input, err := strconv.Atoi(args[1]); err == nil{
        DeleteFromFile(input)
        } else {
            fmt.Println("Please put the index of the todo youre trying to delete")
        }
    default:
    PrintTodos()
    }

    flag.Parse()
}

func PrintTodos(){
	todos := ReadTodosFromFile()
	for i := 0; i < len(todos) - 1; i++{
		fmt.Println(i," "+todos[i])
	}
}

func DeleteFromFile(index int){
	todos := ReadTodosFromFile()
	todos = RemoveIndexFromArray(todos, index)
	WriteToFile(BuildStringFromArray(todos))
}
