package pokebattle 

import(
	"github.com/rgarcia2304/pokedexcli/internal"
	"fmt"
	"math/rand"
	
)

type Battle struct{
	ChallengerName string
	OpponentName string
	Challenger pokeapi.Pokemon 
	Opponent pokeapi.Pokemon
}

func CalculateDamage(pokemonAttacker, pokemonDefender pokeapi.Pokemon) (float64){
	
	//cases for different pokemon and pokemon attack name 
	attackerType := pokemonAttacker.Types[0].Type.Name 
	defenderType := pokemonDefender.Types[0].Type.Name

	//run different pokemon attack scenarios
	switch attackerType {

    	case "normal":
        	switch defenderType {
        	case "rock", "steel":
            		return 0.5
        	case "ghost":
            		return 0.0
        	default:
            		return 1.0
        	}

    case "fire":
        	switch defenderType {
        	case "fire", "water", "rock", "dragon":
            		return 0.5
        	case "grass", "ice", "bug", "steel":
            		return 2.0
        	default:
            		return 1.0
        	}

    	case "water":
        	switch defenderType {
        		case "water", "grass", "dragon":
            			return 0.5
        		case "fire", "ground", "rock":
            			return 2.0
        		default:
            			return 1.0
        	}

    	case "electric":
        	switch defenderType {
        	case "electric", "grass", "dragon":
            		return 0.5
        	case "water", "flying":
            		return 2.0
        	case "ground":
            		return 0.0
        	default:
            		return 1.0
        	}

    	case "grass":
        	switch defenderType {
        	case "fire", "grass", "poison", "flying", "bug", "dragon", "steel":
            		return 0.5
        	case "water", "ground", "rock":
            		return 2.0
        	default:
            		return 1.0
        	}

    	case "ice":
        	switch defenderType {
        	case "fire", "water", "ice", "steel":
            		return 0.5
        	case "grass", "ground", "flying", "dragon":
            		return 2.0
        	default:
            		return 1.0
        	}

    	case "fighting":
        	switch defenderType {
        	case "poison", "flying", "psychic", "bug", "fairy":
            		return 0.5
        	case "normal", "ice", "rock", "dark", "steel":
            		return 2.0
        	case "ghost":
            		return 0.0
        	default:
            		return 1.0
        	}

    	case "poison":
        	switch defenderType {
        	case "poison", "ground", "rock", "ghost":
           		return 0.5
        	case "grass", "fairy":
            		return 2.0
        	case "steel":
            		return 0.0
        	default:
            		return 1.0
        	}

    	case "ground":
        	switch defenderType {
        	case "grass", "bug":
            		return 0.5
        	case "fire", "electric", "poison", "rock", "steel":
            		return 2.0
        	case "flying":
            		return 0.0
        	default:
            		return 1.0
        	}

    	case "flying":
        	switch defenderType {
        	case "electric", "rock", "steel":
            		return 0.5
        	case "grass", "fighting", "bug":
            		return 2.0
        	default:
            		return 1.0
        	}

    	case "psychic":
        	switch defenderType {
        	case "psychic", "steel":
            		return 0.5
        	case "fighting", "poison":
            		return 2.0
        	case "dark":
            		return 0.0
        	default:
            		return 1.0
        	}

    	case "bug":
        	switch defenderType {
        		case "fire", "fighting", "poison", "flying", "ghost", "steel", "fairy":
            			return 0.5
        	case "grass", "psychic", "dark":
            		return 2.0
        	default:
            		return 1.0
        	}

    	case "rock":
        	switch defenderType {
        	case "fighting", "ground", "steel":
            		return 0.5
        	case "fire", "ice", "flying", "bug":
            		return 2.0
        	default:
            		return 1.0
        	}

    	case "ghost":
        	switch defenderType {
        	case "dark":
            		return 0.5
        	case "psychic", "ghost":
            		return 2.0
        	case "normal":
            		return 0.0
        	default:
            		return 1.0
        	}

    	case "dragon":
        	switch defenderType {
        	case "steel":
            		return 0.5
        	case "dragon":
            		return 2.0
        	case "fairy":
            		return 0.0
        	default:
            		return 1.0
        	}

    	case "dark":
        	switch defenderType {
        		case "fighting", "dark", "fairy":
            			return 0.5
        		case "psychic", "ghost":
            			return 2.0
        		default:
            			return 1.0
        	}

    	case "steel":
        	switch defenderType {
        		case "fire", "water", "electric", "steel":
            			return 0.5
        		case "ice", "rock", "fairy":
            			return 2.0
        		default:
            			return 1.0
        	}

    	case "fairy":
        	switch defenderType {
        	case "fire", "poison", "steel":
            		return 0.5
        	case "fighting", "dragon", "dark":
            		return 2.0
        	default:
            		return 1.0
        	}
    }

    return 1.0	
}
func (b *Battle) SimulateBattle() (string){

	if b == nil{
		return fmt.Sprintf("There are no opponents")
	}

	turn := 0  
	challengerHp := float64(b.Challenger.Stats[0].BaseStat)
	opponentHp := float64(b.Opponent.Stats[0].BaseStat)
		//decide who goes first in the battle, decided by speed
	if b.Challenger.Stats[5].BaseStat > b.Opponent.Stats[5].BaseStat{
		turn = 0 
	}else{
		turn = 1
	}

	for{
		randomNum := rand.Intn(10)
		randomMultiplier := float64(randomNum) * 0.1
		switch turn{
			case 0: //challenger is attacking
				fmt.Print(b.ChallengerName + " is attacking \r\n")

				damageFromAttack := (float64(b.Challenger.Stats[1].BaseStat) / float64(b.Opponent.Stats[2].BaseStat)) * float64(b.Challenger.Stats[1].BaseStat) * CalculateDamage(b.Challenger, b.Opponent) * randomMultiplier
				fmt.Print(damageFromAttack)
				fmt.Print("\r\n")
				attackMsg := fmt.Sprintf("%v just sent an attack with %v \r\n", b.ChallengerName, damageFromAttack)
				fmt.Print(attackMsg)
				opponentHp -= damageFromAttack 
			case 1: // oponent is attacking
				fmt.Print(b.OpponentName + " is attacking \r\n")
				damageFromAttack := (float64(b.Opponent.Stats[1].BaseStat) / float64(b.Challenger.Stats[2].BaseStat)) * float64(b.Opponent.Stats[1].BaseStat) * CalculateDamage(b.Opponent, b.Challenger) * randomMultiplier
				attackMsg := fmt.Sprintf("%v just sent an attack with %v \r\n", b.OpponentName, damageFromAttack)	
				fmt.Print(float64(b.Opponent.Stats[1].BaseStat / b.Challenger.Stats[2].BaseStat))

				fmt.Print("\r\n")
				fmt.Print(float64(b.Challenger.Stats[2].BaseStat))

				fmt.Print("\r\n")
				fmt.Print(float64(b.Opponent.Stats[1].BaseStat))

				fmt.Print("\r\n")
				fmt.Print(CalculateDamage(b.Opponent, b.Challenger))
				fmt.Print("\r\n")
				fmt.Print(randomMultiplier)
				fmt.Print("\r\n")

				fmt.Print(attackMsg)
				challengerHp -= damageFromAttack
			default:
				return fmt.Sprintf("This should never happen \r\n")
		}
		
		challengerHpMsg := fmt.Sprintf("%v Hp: %v \r\n",b.ChallengerName, challengerHp)
		fmt.Print(challengerHpMsg)
		
		opponentHpMsg := fmt.Sprintf("%v Hp: %v \r\n",b.OpponentName, opponentHp)
		fmt.Print(opponentHpMsg)

		if turn == 0{
			turn = 1
		}else{
			turn = 0
		}

		if challengerHp <= 0{
			return fmt.Sprintf("%v has won \r\n", b.OpponentName)
		}else if opponentHp <= 0{
			return fmt.Sprintf("%v has won \r\n", b.ChallengerName)
		}

	}




}
