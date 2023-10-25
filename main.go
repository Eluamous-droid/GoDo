package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/eluamous-droid/godo/remote"
	"github.com/eluamous-droid/godo/todo/tools"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printTodos()
	} else {
		switch args[0] {
		case "add":
			if args[1] != "" {
				tools.AppendToFile(args[1])
			} else {
				fmt.Println("Missing todo to add")
			}
		case "rm":
			if input, err := strconv.Atoi(args[1]); err == nil {
				deleteFromFile(input)
			} else {
				fmt.Println("Please put the index of the todo youre trying to delete")
			}
		case "remote":
			var items []remote.TodoItem
			items = remote.GetAllItemsInGroup("work")
			for _, item := range items {
				println("Group: " + item.Group + " Task: " + item.Task)
			}
		default:
			fmt.Println("Please use a valid command")
		}
	}

	flag.Parse()
}
func printTodos() {
	todos := tools.ReadTodosFromFile()
	for i := 0; i < len(todos)-1; i++ {
		fmt.Println(i, " "+todos[i])
	}
}

func deleteFromFile(index int) {
	todos := tools.ReadTodosFromFile()
	todos = tools.RemoveIndexFromArray(todos, index)
	tools.WriteToFile(tools.BuildStringFromArray(todos))
}
