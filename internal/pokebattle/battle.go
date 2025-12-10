package pokebattle 

import(
	"github.com/rgarcia2304/pokedexcli/internal"
	"fmt"
	
)

type Battle struct{
	ChallengerName string
	OpponentName string
	Challenger Pokemon 
	Opponent Pokemon
}


func (b *Battle) simulateBattle() (string){

	if Challenger == nil || Opponent == nil{
		return fmt.Sprintf("There are no opponents")
	}

	turn := 0 
	challengerHp := b.Challenger.Stats[0].BaseStat
	opponentHp := b.Opponent.Stats[0].BaseStat
	
	//decide who goes first in the battle, decided by speed
	if b.Challenger.Stats[5].BaseStat > b.Opponent.Stats[5].BaseStat{
		turn = 0 
	}else{
		turn = 1
	}

	for{
		switch turn{
			case turn == 0: //challenger is attacknig 
				damageFromAttack := float(b.Challenger.Stats[1].BaseStat / b.Opponent.Stats[2].BaseStat) * float(b.Challenger.Stats[1].BaseStat
				opponentHp -= damageFromAttack 
			case turn == 1: // oponent is attacking
				damageFromAttack := float(b.Opponent.Stats[1].BaseStat / b.Challenger.Stats[2].BaseStat) * float(b.Opponent.Stats[1].BaseStat
				challengerHp -= damageFromAttack
			default:
				return fmt.Sprintf("This should never happen \r\n")
		}

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
