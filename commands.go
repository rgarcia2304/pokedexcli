package main

import(

	"os"
	"fmt"
	"errors"
)


type cliCommand struct{
	name string 
	description string
	callback func(*Config, ...string) error
}




func commandExit(cfg *Config, args ...string) error{
	fmt.Println("Closing the Pokedex... Gooddbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args ...string) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	for key, val := range validCommands{
		help_menu := fmt.Sprintf("%v: %v", key, val.description)
		fmt.Println(help_menu)
	}
	return nil
}

func commandMapf(cfg *Config, args ...string) error{
	
	//get the initial response

	locations, err := cfg.Client.ListLocations(*cfg.nextURL)

	if cfg.nextURL == nil{
		return errors.New("This is the last page")
	}

	if err != nil{
		return errors.New("Issue with the API fetching resource")
	}

	for _, location:= range locations.Results{
			fmt.Println(location.Name)
		}
	
	
	// update the previous and new based on the repsonse
	cfg.nextURL = locations.Next
	cfg.previousURL = locations.Previous
	return nil
}

func commandMapb(cfg *Config, args ...string) error{
	
	if cfg.previousURL == nil{
		return errors.New("You are on the first page")
	}

	locations, err := cfg.Client.ListLocations(*cfg.previousURL)
	if err != nil{
		return errors.New("Issue with fetching API")
	}

	for _, location := range locations.Results{
		fmt.Println(location.Name)
		}
	
	
	// update the previous and new based on the repsonse
	cfg.nextURL = locations.Next
	cfg.previousURL = locations.Previous
	return nil

}

func commandExplore(cfg *Config, args ...string) error{
	
	//build full url
	if len(args) == 0{
		fmt.Println("You must proved a location")
	}
	areaName := args[0]

	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	exploreUrl := baseUrl + areaName
	pokeNames, err := cfg.Client.ExploreArea(exploreUrl)
	if err != nil{
		return errors.New("Issue with fetching API")
	}

	for _, pokemonName := range pokeNames.PokemonEncounters{
		fmt.Println(pokemonName)
		}
	
		return nil
}

