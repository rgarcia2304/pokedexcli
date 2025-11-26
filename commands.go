package main

import(

	"os"
	"fmt"
	"github.com/rgarcia2304/pokedecxcli/internal/pokeapi"
)


type cliCommand struct{
	name string 
	description string
	callback func() error
}




func commandExit(*config) error{
	fmt.Println("Closing the Pokedex... Gooddbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*Config) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	for key, val := range validCommands{
		help_menu := fmt.Sprintf("%v: %v", key, val.description)
		fmt.Println(help_menu)
	}
	return nil
}

func commandMapf(*Config) error{
	
	//get the initial response 
	locations,err := Config.pokeapi.Client.ListLocations(*config.nextURL)

	if locations.Next == nil{
		fmt.Println("You are on the first page")	
	}else{
		for _, location:= range location{
			fmt.Println(location.Name)
		}
	}
	
	// update the previous and new based on the repsonse
	Config.nextURL = locations.Next
	Config.prevURL = locations.Prev
	return nil
}

func commandMapb(*Config) error{
	locations, err := Config.pokeapi.Client.ListLocations(*config.nextURL)
	if locations.Prev == nil{
		fmt.Println("You are on the first page")
	}else{
		for _, location := range locations.Result{
			fmt.Println(location.Name)
		}
	}
	
	// update the previous and new based on the repsonse
	Config.nextURL = locations.Next
	Config.prevURL = locations.Prev
	return nil

}

