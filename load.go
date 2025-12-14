package main 

import(

	"errors"
	"internal/pokesave"
	"internal/pokeapi"
)


func loadSave(cfg *Config) (error){

	if err := pokeSave.Load("./file.tmp", cfg.pokeDeck); err != nil{
		return errors.New("Failed to load saved progress")
	}
	
	return nil
}
