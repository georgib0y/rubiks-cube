package main

import (
	"fmt"
	"strings"
)

var faceAdjacencies map[Colour][]Colour = map[Colour][]Colour{
	Yellow: {Green, Red, Blue, Orange},
	White:  {Red, Green, Orange, Blue},
	Blue:   {Yellow, Red, White, Orange},
	Green:  {Red, Yellow, Orange, White},
	Orange: {Green, Yellow, Blue, White},
	Red:    {White, Blue, Yellow, Green},
}

type Cube struct {
	faces map[Colour]*Face
	front *Face
}

func NewCube() Cube {
	cube := Cube{
		faces: map[Colour]*Face{},
	}

	for _, colour := range colours {
		face := NewFace(colour)
		cube.faces[colour] = &face
	}

	for colour, adjs := range faceAdjacencies {
		for _, adj := range adjs {
			cube.faces[colour].AddAdjacentFace(cube.faces[adj])
		}
	}

	cube.front = cube.faces[Red]

	return cube
}

func (c *Cube) Show() {
	c.faces[White].OrientTop(Orange)
	white := strings.Split(c.faces[White].String(), "\n")

	c.faces[Green].OrientTop(White)
	green := strings.Split(c.faces[Green].String(), "\n")

	c.faces[Red].OrientTop(White)
	red := strings.Split(c.faces[Red].String(), "\n")

	c.faces[Yellow].OrientTop(Red)
	yellow := strings.Split(c.faces[Yellow].String(), "\n")

	c.faces[Blue].OrientTop(White)
	blue := strings.Split(c.faces[Blue].String(), "\n")

	c.faces[Orange].OrientTop(White)
	orange := strings.Split(c.faces[Orange].String(), "\n")

	s := ""
	for i := 0; i < 3; i++ {
		s += fmt.Sprintf("      %s\n", white[i])
	}

	for i := 0; i < 3; i++ {
		s += fmt.Sprintf("%s %s %s %s\n", green[i], red[i], blue[i], orange[i])
	}

	for i := 0; i < 3; i++ {
		s += fmt.Sprintf("      %s\n", yellow[i])
	}

	fmt.Println(s)
}

func (c *Cube) Orient(front, top Colour) {
	c.front = c.faces[front]
	c.front.OrientTop(top)
}

func (c *Cube) Up() {
	topColour := c.front.TopFace().colour
	toSwap := c.front.TopSquares()
	face := c.front.LeftFace()
	for i := 0; i < 4; i++ {
		face.OrientTop(topColour)
		toSwap = face.SwapSquares(0, toSwap)
		face = face.LeftFace()
	}
}

func (c *Cube) Down() {
	topColour := c.front.TopFace().colour
	c.Orient(c.front.colour, c.front.BottomFace().colour)
	c.Up()
	c.Orient(c.front.colour, topColour)
}

func (c *Cube) Right() {
	topColour := c.front.TopFace().colour
	c.Orient(c.front.colour, c.front.RightFace().colour)
	c.Up()
	c.Orient(c.front.colour, topColour)
}

func (c *Cube) Left() {
	topColour := c.front.TopFace().colour
	c.Orient(c.front.colour, c.front.LeftFace().colour)
	c.Up()
	c.Orient(c.front.colour, topColour)
}

func (c *Cube) Front() {
	topColour := c.front.TopFace().colour
	frontColour := c.front.colour
	c.Orient(topColour, frontColour)
	c.Up()
	c.Orient(frontColour, topColour)
}

func (c *Cube) Back() {
	topColour := c.front.TopFace().colour
	frontColour := c.front.colour
	c.Orient(topColour, frontColour)
	c.Down()
	c.Orient(frontColour, topColour)
}
