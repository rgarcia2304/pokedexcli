package main 


import(

"fmt"
"bufio"
"os"
"github.com/rgarcia2304/pokedexcli/internal"
"github.com/rgarcia2304/pokedexcli/internal/pokecache"
"github.com/rgarcia2304/pokedexcli/internal/pokebattle"
"golang.org/x/term"
"time"
"log"
)

var validCommands map[string]cliCommand 

type Config struct{
	nextURL *string
	previousURL *string
	pokeapi.Client 
	pokecache.Cache
	pokeDeck map[string]pokeapi.Pokemon
	pokebattle.Battle
}

type REPL struct{
	history []string
	historyIndex int
	buffer []rune
	hasDraft bool
	draftBeforeHistory string
	prompt string
	
}


func main(){
	
	func() {
		validCommands = map[string]cliCommand{
    		"exit": {
        		name:        "exit",
        		description: "Exit the Pokedex",
        		callback:    commandExit,
    		},
		"help":{
			name:       "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"mapf":{
			name: "map",
			description: "Gives locations of pokemon",
			callback: commandMapf,
		},
		"mapb":{
			name:"mapb",
			description: "Gives locations of pokemon",
			callback: commandMapb,
		},
		"explore":{
			name:"explore",
			description: "Find out what pokemon are in the area",
			callback: commandExplore,
		},
		"catch":{
			name: "catch",
			description: "Performs a catching action on a pokemon",
			callback: commandCatch, 
		},
		"inspect":{
			name: "inspect",
			description: "View info about pokemon that you have caught",
			callback: commandInspect, 
		},
		"pokedeck":{
			name: "pokedeck", 
			description: "View all the pokemon that you have caught",
			callback: commandPokedeck,
			
		},
		"battle":{
			name: "battle", 
			description: "Include two pokemon in your deck to battle out with, provide both of them as command line arguments",
			callback: commandBattle,
			
		},
		"save":{
			name: "save", 
			description: "allows you to save pokemon in your deck between sessions",
			callback: commandSave,
			
		},
	}			
	
	}()
	
	//initialize the cache that will be used 
	pokemonCache := pokecache.NewCache(time.Second * 5)
	client := pokeapi.NewClient(time.Second * 10, pokemonCache)
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	deck := make(map[string]pokeapi.Pokemon)
	
	//load the saved pokeDeck
	init_config := Config{nextURL: &baseURL, Client: client, pokeDeck: deck}
	//load the saved pokedeck 
	loadSave(init_config)

	terminalREPL := REPL{historyIndex: 0, prompt: "Pokedeck> " }	
	
		//put terminal in raw mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil{
		log.Fatalf("failed to put terminal in raw mode: %v", err)
	}
	
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(terminalREPL.prompt)
	for{
		//if an arrow command was given then have it put that command in the scanner 

		b, err := reader.ReadByte()
		if err != nil{
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		}

		switch b{

		case '\r', '\n':
			handleEnter(&init_config, &terminalREPL)
		
		case 127:
			handleBackSpace(&terminalREPL)

		case 27:
			seq1, _ := reader.ReadByte()
			seq2, _ := reader.ReadByte()
			if seq1 == '['{
				switch seq2{
				case 'A':
					handleUp(&terminalREPL)
				case 'B': 
					handleDown(&terminalREPL)
				}
			}
		default:
			if b >=32 && b < 127{
				terminalREPL.buffer = append(terminalREPL.buffer, rune(b))
				fmt.Printf("%c", b)
			}
		}

		}
	
}
