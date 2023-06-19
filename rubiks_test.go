package main

import (
	"fmt"
	"testing"
)

func TestOrientFace(t *testing.T) {
	face := NewFace(Red)

	for i := 0; i < 4; i++ {
		adj := NewFace(Colour(i))
		face.AddAdjacentFace(&adj)
	}

	for i := 0; i < 8; i++ {
		face.squares[i] = Colour(i % 6)
	}

	fmt.Println(face.String())
	face.OrientTop(White)
	fmt.Println(face.String())
}

func TestSwapSquares(t *testing.T) {
	face := NewFace(Red)

	for i := 0; i < 8; i++ {
		face.squares[i] = Colour(i % 6)
	}

	fmt.Println(face.String())
	swap := []Colour{Blue, Blue, Blue}

	swap = face.SwapSquares(0, swap)

	fmt.Println(face.String())
	fmt.Println(swap)
}

func TestShowCube(t *testing.T) {
	cube := NewCube()

	cube.Show()
	cube.Orient(Red, White)
	cube.Left()
	cube.Show()
}
