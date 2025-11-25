package main 


import(

"fmt"
"bufio"
"os"
)


func main(){
	scanner := bufio.NewScanner(os.Stdin)

	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scannerVal := scanner.Text()
		//clean the scanned value
		cleanedScan := cleanInput(scannerVal)

		//capture first word of input 
		term_msg := fmt.Sprintf("Your command was: %v", cleanedScan[0])
		fmt.Println(term_msg)
	}

}
