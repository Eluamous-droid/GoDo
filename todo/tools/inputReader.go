package tools

import (
	"flag"
	"os"
)


func ReadInput(){

    args := os.Args[1:]
	//todoPtr := flag.String("ToDo", "", "Text to save.")
    //listCommand := flag.NewFlagSet("list", flag.ExitOnError)

    switch args[0] {
    case "list":
        ReadTodos()
    default:
    SaveToFile(args[0])
    }

    flag.Parse()

    //if *todoPtr == "" {
      //  flag.PrintDefaults()
        //os.Exit(1)
		
    //}
    
    //SaveToFile(args[0])
	
}