package main 


import(

"fmt"
"bufio"
"os"
"github.com/rgarcia2304/pokedexcli/internal"
"github.com/rgarcia2304/pokedexcli/internal/pokecache"

"time"
)

var validCommands map[string]cliCommand 

type Config struct{
	nextURL *string
	previousURL *string
	pokeapi.Client 
	pokecache.Cache
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
	}			
	
	}()

	//initialize all the config struct fields 
	
	//initialize the cache that will be used 
	pokemonCache := pokecache.NewCache(time.Second * 5)
	client := pokeapi.NewClient(time.Second * 10, pokemonCache)
	baseURL := "https://pokeapi.co/api/v2/location-area/" 
	
	init_config := Config{nextURL: &baseURL, Client: client}
	
	
	//start scanning for input 
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scannerVal := scanner.Text()
		//clean the scanned value
		cleanedScan := cleanInput(scannerVal)

		//capture first word of input
		cmd, ok := validCommands[cleanedScan[0]]
		
		if !ok{
			fmt.Println("Unknown Command")
			continue
		}else{

			var args []string
			if len(cleanedScan) > 1{
				args = cleanedScan[1:]
			}

			//act on the action of the command
			if err := cmd.callback(&init_config, args...); err != nil{
				fmt.Println(err)
			}
		}
	}

}
