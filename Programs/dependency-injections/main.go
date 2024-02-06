package main

import (
	"fmt"


)

// Toy is a struct representing a toy
type Toy struct {
	Name  string
	Color string
}

// Box is a struct representing a box of toys
type Box struct {
	Contents *Toy // Contents is a pointer to a Toy
}

// Playable is an interface for toys that can be played with
type Playable interface {
	Play()
}

// Implementing the Play method for Toy
func (t *Toy) Play() {
	fmt.Printf("Playing with a %s %s\n", t.Color, t.Name)
}

func main() {
	// Create a new toy
	actionFigure := &Toy{
		Name:  "Action Figure",
		Color: "Red",
	}

	// Create a box and put the toy inside
	toyBox := &Box{
		Contents: actionFigure,
	}

	// Access the toy inside the box using the pointer
	fmt.Printf("The toy in the box is a %s %s\n", toyBox.Contents.Color, toyBox.Contents.Name)

	// Play with the toy using the interface
	playWithToy(actionFigure)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

}
// Function to play with a toy using the Playable interface
func playWithToy(p Playable) {
	p.Play()
}
