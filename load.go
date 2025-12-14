package main 

import(

"errors"
"github.com/rgarcia2304/pokedexcli/internal/pokesave"
)


func loadSave(cfg *Config) (error){

	if err := pokesave.Load("./file.tmp", cfg.pokeDeck); err != nil{
		return errors.New("Failed to load saved progress")
	}
	
	return nil
}
