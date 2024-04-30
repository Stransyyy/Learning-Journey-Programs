package main

import "fmt"

// interfaces example

// ABSTRACT TYPE --
//
//	|
//	|
//	V
type Player interface {
	KickBall() int
}

// CONCRETE TYPE --
//                |
//                |
//                V
// type Player struct {
// 	Power int
// }

// CONCRETE TYPE TOO
type BadPlayer struct {
	Power int
}

func (badPlayer BadPlayer) KickBall() int {
	// This exampple will always execute this logic
	//1.
	//2.
	//3.

	fmt.Println("BAD player is kicking the ball with ", badPlayer.Power)

	return badPlayer.Power
}

/*This was part of the struct Player that is now commented out
func (p Player) KickBall() {
	// This exampple will always execute this logic
	//1.
	//2.
	//3.

	fmt.Println("Player is kicking the ball with power ", p.Power)
}
*/

type GoodPlayer struct {
	Power int
}

func (goodPlayer GoodPlayer) KickBall() int {
	multiplier := 2

	fmt.Println("GOOD player is kicking the ball with ", goodPlayer.Power*multiplier)

	return goodPlayer.Power
}

func main() {
	/* Used in the first example of only using strcuts and not interfaces
	player := Player{
		Power: 100,
	}

	betterPlayer := Player{
		Power: 200,
	}

	TeamPlayerKickBall(player)
	TeamPlayerKickBall(betterPlayer)
	*/

	badPlayer := BadPlayer{
		Power: 50,
	}

	GoodPlayer := GoodPlayer{
		Power: 100,
	}

	TeamPlayerKickBall(badPlayer) //---------> error here: cannot use badPlayer (variable of type BadPlayer) as Player value in argument to TeamPlayerKickBall: BadPlayer does not implement Player (wrong type for method KickBall) have KickBall() want KickBall() int
	//This is beacuse the function KickBall() Does not return an integer like specified in the type interface Player
	// If we modify the KickBall() --Line 29-- function to return an int it will work now

	TeamPlayerKickBall(GoodPlayer)
}

func TeamPlayerKickBall(player Player) {
	player.KickBall()
}
