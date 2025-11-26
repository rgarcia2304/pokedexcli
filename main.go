package main 


import(

"fmt"
"bufio"
"os"
)

var validCommands map[string]cliCommand 

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
	}			
	
	}()
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
