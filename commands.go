package main

import(

	"os"
	"fmt"
	"errors"
	"math/rand"
)


type cliCommand struct{
	name string 
	description string
	callback func(*Config, ...string) error
}




func commandExit(cfg *Config, args ...string) error{
	fmt.Print("Closing the Pokedex... Gooddbye!")
	fmt.Print("\r\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args ...string) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("\r\n")
	fmt.Println("Usage")
	fmt.Print("\r\n")
	for key, val := range validCommands{
		help_menu := fmt.Sprintf("%v: %v", key, val.description)
		fmt.Print(help_menu)
		fmt.Print("\r\n")
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
			fmt.Print(location.Name)
			fmt.Print("\r\n")
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
		fmt.Print(location.Name)
		fmt.Print("\r\n")
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
		fmt.Print(pokemonName.Pokemon.Name)
		fmt.Print("\r\n")
		}
	
	return nil
}

func commandCatch(cfg *Config, args ...string) error{
	
	//build full url
	if len(args) == 0{
		fmt.Println("You must provide a pokemon to catch")
	}
	pokemonName := args[0]

	//before even hitting cache if pokemon is caught don't let them catch more
	_, ok := cfg.pokeDeck[pokemonName]
	if ok{
		return errors.New("Already have caught this pokemon")
	}

	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	pokeUrl := baseUrl + pokemonName
	pokeResp, err := cfg.Client.CatchPokemon(pokeUrl)
	
	if err != nil{
		return errors.New("Issue with fetching API")
	}
	
	val := fmt.Sprintf("This is the pokemon xp %v", pokeResp.BaseExperience)
	fmt.Print(val)
	fmt.Print("\r\n")

	//process whether it can be caught or not 
        pokeBallNumber := rand.Intn(150) 
	fmt.Print("Throwing a ball at " + pokemonName)
	fmt.Print("\r\n")
	
	if pokeBallNumber >= pokeResp.BaseExperience{
		//add this to the pokemon deck 
		cfg.pokeDeck[pokemonName] = pokeResp
		fmt.Print(pokemonName + " was caught!")
		fmt.Print("\r\n")
	}else{
		fmt.Println(pokemonName + " escaped")
		fmt.Print("\r\n")
	}

	return nil
}

func commandInspect(cfg *Config, args ...string) error{
	
	//build full url
	if len(args) == 0{
		return errors.New("You must provide a pokemon to catch")
	}
	pokemonName := args[0]

	//before even hitting cache if pokemon is caught don't let them catch more
	pokeInfo, ok := cfg.pokeDeck[pokemonName]
	if !ok{
		return errors.New("You do not have this Pokemon")
	}

	//Now loop through the stats and format 
	fmt.Print("Name: " + pokemonName)
	fmt.Print("\r\n")
	
	heightValue := fmt.Sprintf("Height: %v", pokeInfo.Height)
	fmt.Print(heightValue)
	fmt.Print("\r\n")

	weightValue := fmt.Sprintf("Weight: %v", pokeInfo.Weight)
	fmt.Print(weightValue)
	fmt.Print("\r\n")

	//now loop over the different stats
	fmt.Print("Stats:")
	fmt.Print("\r\n")
	for _, val := range pokeInfo.Stats{
		res := fmt.Sprintf("   -%v: %v", val.Stat.Name, val.BaseStat)
		fmt.Print(res)
		fmt.Print("\r\n")
	}

	fmt.Print("Types:")
	fmt.Print("\r\n")
	for _, val := range pokeInfo.Types{
		res := fmt.Sprintf("   -%v", val.Type.Name)
		fmt.Print(res)
		fmt.Print("\r\n")
	}
		
	return nil
}

func commandPokedeck(cfg *Config, args ...string) (error){
	if len(args) > 0{
		return errors.New("Invalid Command Sequence")
	}

	if len(cfg.pokeDeck) == 0 {
		return errors.New("Your Pokedeck is empty")
	}
	
	fmt.Print("Your Pokemon:")
	fmt.Print("\r\n")
	for pokemon, _:= range cfg.pokeDeck{
		fmt.Print("- " + pokemon)
		fmt.Print("\r\n")
	}
	
	return nil
}

func commandBattle(cfg *Config, args ...string) (error){
	if len(args) != 2 {
		return errors.New("You must include two pokemon to battle")
	}

	challengerPokeStats, found := cfg.pokeDeck[args[0]]
	
	if !found{
		return errors.New("Please select a challenger pokemon in your deck to battle with")
	}

	opponentPokeStats, found := cfg.pokeDeck[args[1]]
	
	if !found{
		return errors.New(" Please select an opponent pokemon in your deck to battle with")
	}
	
	cfg.Battle.OpponentName = args[1]
	cfg.Battle.ChallengerName = args[0]
	cfg.Battle.Challenger = challengerPokeStats
	cfg.Battle.Opponent = challengerPokeStats

	battleResult := cfg.Battle.simulateBattle()
	fmt.Print(battleResult)
	
}
