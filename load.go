package main 

import(

"errors"
"internal/pokesave"
"github.com/rgarcia2304/pokedexcli/internal"
"github.com/rgarcia2304/pokedexcli/internal/pokesave"
)


func loadSave(cfg *Config) (error){

	if err := pokeSave.Load("./file.tmp", cfg.pokeDeck); err != nil{
		return errors.New("Failed to load saved progress")
	}
	
	return nil
}
