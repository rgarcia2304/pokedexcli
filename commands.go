package main

import(

	"os"
	"fmt"
)


type cliCommand struct{
	name string 
	description string
	callback func() error
}


func commandExit()error{
	fmt.Println("Closing the Pokedex... Gooddbye!")
	os.Exit(0)
	return nil
}

func commandHelp()error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	for key, val := range validCommands{
		help_menu := fmt.Sprintf("%v: %v", key, val.description)
		fmt.Println(help_menu)
	}
	return nil
}


