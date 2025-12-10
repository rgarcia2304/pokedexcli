package main

import (
    "fmt"
    "unicode/utf8"
    "strings"
)



func handleEnter(c *Config, r *REPL){


	fmt.Print("\r\n")

	line := string(r.buffer)
	trimmed := trimSpace(line)

	parsedInput := strings.Fields(trimmed)

	//perform callback upon 
	cmd, ok := validCommands[parsedInput[0]]

	if !ok{
		fmt.Println("Unknown Command")
	}else{
		var args []string
		if len(parsedInput) > 1{
			args = parsedInput[1:]
		}
		if err := cmd.callback(c, args...); err != nil{
			fmt.Println(err)
		}
	}

	//add command to history
	r.history = append(r.history, trimmed)
	
	r.buffer = r.buffer[:0]
	r.historyIndex = len(r.history)
	r.hasDraft = false
	r.draftBeforeHistory = ""
	fmt.Print(r.prompt)
	renderLine(r)
	}

func handleBackSpace(r *REPL){
	if len(r.buffer) == 0{
		return
	}

	r.buffer = r.buffer[:len(r.buffer)-1]
	renderLine(r)
}

func handleUp(r *REPL){
	if len(r.history) == 0{
		return
	}

	if r.historyIndex == len(r.history) && !r.hasDraft{
		r.draftBeforeHistory = string(r.buffer)
		r.hasDraft = true
	}
	if r.historyIndex > 0{
		r.historyIndex--
	}

	r.buffer = []rune(r.history[r.historyIndex])
	renderLine(r)
}

func handleDown(r *REPL){
	if len(r.history) == 0{
		return 
	}

	if r.historyIndex < len(r.history)-1{
		
		r.historyIndex++ 
		r.buffer = []rune(r.history[r.historyIndex])
	}else if r.historyIndex == len(r.history)-1{
		//move back to whatever was drafted or the new line
		r.historyIndex = len(r.history)
		if r.hasDraft{
			r.buffer = []rune(r.draftBeforeHistory)
		}else {
			r.buffer = r.buffer[:0]
		}
		
		r.hasDraft = false
		r.draftBeforeHistory = ""
	}else{
		return //case for being on a newline
	}

	renderLine(r)
}

func renderLine(r *REPL){

	fmt.Print("\r") // carriage return
	fmt.Print("\x1b[2K") //prints new line
	fmt.Print(r.prompt)
	fmt.Print(string(r.buffer))
}

func trimSpace(s string) string{
	for len(s) > 0 {
        r, size := utf8.DecodeRuneInString(s)
        if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
            s = s[size:]
        } else {
            break
        }
    }
    for len(s) > 0 {
        r, size := utf8.DecodeLastRuneInString(s)
        if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
            s = s[:len(s)-size]
        } else {
            break
        }
    }
    return s

}


