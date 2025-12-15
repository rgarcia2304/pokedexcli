# Pokedeck command line interface application 

## Requirements
Make sure you have an installation of golang active on your machine 

Simply clone this repository into your local machine 

## In your command line run the following commands 

### Build the project
 
 `go bulid`

### Update the modules required for this project

`go mod tidy`

### To run the project 

 `go run .` or  `./pokedexcli`


## Usage 

Use `help` to get the different commands available.


### Some of the most useful commands 


Get history of your commands

` Up and Down Arrow`

Get the locations where different pokemon could be 

`mapf`

Explore the different locations where pokemon could be 

`explore (location you want to explore)`

## Architecture 

### REPL 

This application is built as a Read Evaluate Print Loop interface.
Users enter commands, which call callback functions which perform a variety of different functions. 

### Terminal 

This application has custom rendering to support UNIX like functionallity which I really like using the x/term library. 
The terminal supports various functionality including support for backspace, entering commands, up and down arrows for commands history, and string cleaning. 

### Commands 

Several commands are supported on this cli application. When submitting a command on the command line interface, a callback is called to call the respective command.
### API design 

Some commands require calling of API's. A pokeapi package was created to provide support for commands which require api calls.
An HTTP client was used to support those calls, with timeouts in case no resources were found at endpoint being hit. 
Upon getting resources they are then unmarshalled and printed to the terminal depending on use case:

For example:

`Pokedeck> mapf` 

Response:
canalave-city-area
eterna-city-area
pastoria-city-area
...

mapf ---> hits the "https://pokeapi.co/api/v2/location-area/ endpoint --> unmarshall resource into location struct --> print to terminal

###  Caching

This application uses a go routine, to update the cache in constant intervals using LRU principles. 

## Persistent Saving 

The save command provides users the ability to save pokemon between sessions to there local disk/ 
After all go to catch them all!

## Note 

Feel free to leave feedback or add to the project, thanks!


