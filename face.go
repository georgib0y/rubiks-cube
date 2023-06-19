package main

import "fmt"

type Colour int

const (
	Yellow Colour = iota
	White
	Blue
	Green
	Orange
	Red
)

var colours = [6]Colour{Yellow, White, Blue, Green, Orange, Red}

func (c Colour) String() string {
	switch c {
	case Yellow:
		return "Yellow"
	case White:
		return "White"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Orange:
		return "Orange"
	case Red:
		return "Red"
	default:
		return "Unknown Colour"
	}
}

func (c Colour) StringShort() string {
	switch c {
	case Yellow:
		return "Y"
	case White:
		return "W"
	case Blue:
		return "B"
	case Green:
		return "G"
	case Orange:
		return "O"
	case Red:
		return "R"
	default:
		return "U"
	}
}

type Face struct {
	colour Colour
	// all adjacent faces orderded clockwise starting from the top face around to the left
	adj []*Face
	// all squares orderd clockwise starting from the top left corner around to the middle left
	squares []Colour
}

func NewFace(colour Colour) Face {
	return Face{
		colour:  colour,
		adj:     []*Face{},
		squares: []Colour{colour, colour, colour, colour, colour, colour, colour, colour},
	}
}

func (f *Face) TopFace() *Face {
	return f.adj[0]
}

func (f *Face) LeftFace() *Face {
	return f.adj[3]
}

func (f *Face) BottomFace() *Face {
	return f.adj[2]
}

func (f *Face) RightFace() *Face {
	return f.adj[1]
}

func (f *Face) TopSquares() []Colour {
	return f.squares[:3]
}

func (f *Face) LeftSquares() []Colour {
	return f.squares[2:5]
}

func (f *Face) BottomSquares() []Colour {
	return f.squares[4:7]
}

func (f *Face) RightSquares() []Colour {
	right := []Colour{}
	right = append(right, f.squares[6:]...)
	right = append(right, f.squares[0])
	return right
}

func (f Face) String() string {
	face := fmt.Sprintf("%s %s %s\n%s(%s)%s\n%s %s %s\n",
		f.squares[0].StringShort(),
		f.squares[1].StringShort(),
		f.squares[2].StringShort(),
		f.squares[3].StringShort(),
		f.colour.StringShort(),
		f.squares[4].StringShort(),
		f.squares[5].StringShort(),
		f.squares[6].StringShort(),
		f.squares[7].StringShort(),
	)

	return face

	// out := fmt.Sprintf("Colour: %s\t", f.colour)

	// out += "Faces: "
	// for _, face := range f.adj {
	// 	out += fmt.Sprintf("%s ", face.colour)
	// }

	// out += "\tSquares: "
	// for _, square := range f.squares {
	// 	out += fmt.Sprintf("%s ", square)
	// }

	// return out
}

func (f *Face) AddAdjacentFace(adj *Face) {
	f.adj = append(f.adj, adj)
}

func (f *Face) OrientTop(topColour Colour) {
	sliceIdx := 0

	for i, adj := range f.adj {
		if adj.colour == topColour {
			sliceIdx = i
			break
		}
	}

	f.rotateAdj(sliceIdx)
	f.rotateSquares(sliceIdx * 2)
}

func (f *Face) rotateAdj(idx int) {
	adj := []*Face{}
	adj = append(adj, f.adj[idx:]...)
	adj = append(adj, f.adj[:idx]...)
	f.adj = adj
}

func (f *Face) rotateSquares(idx int) {
	squares := []Colour{}
	squares = append(squares, f.squares[idx:]...)
	squares = append(squares, f.squares[:idx]...)
	f.squares = squares
}

func (f *Face) SwapSquares(start int, toSwap []Colour) []Colour {
	end := start + 3

	swapped := f.squares[start:end]

	squares := []Colour{}
	squares = append(squares, f.squares[:start]...)
	squares = append(squares, toSwap...)
	squares = append(squares, f.squares[end:]...)

	f.squares = squares
	return swapped
}
