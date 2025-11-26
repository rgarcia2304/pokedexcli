package main 


import(

"fmt"
"bufio"
"os"
"internal/pokeapi"
"time"
)

var validCommands map[string]cliCommand 

type Config struct{
	nextURL* string
	previousURL* string
	pokeapi.Client 
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
			callback: mapf,
		},
		"mapb":{
			name:"mapb",
			description: "Gives locations of pokemon",
			callback: mapb,
		},
	}			
	
	}()

	//initialize all the config struct fields 

	client := NewClient(time.Second * 10)
	
	init_config := Config{
		nextURL: "https://pokeapi.co/api/v2/location-area/",
		pokeapi.Client: client}
	
	
	//start scanning for input 
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scannerVal := scanner.Text()
		//clean the scanned value
		cleanedScan := cleanInput(scannerVal)

		//capture first word of input
		val, err := validCommands[cleanedScan[0]]
		if !err{
			fmt.Println("Unknown Command")
		}else{
			//act on the action of the command 
			val.callback()		
		}

		
	}

}
