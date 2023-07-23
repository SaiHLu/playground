package main

type Color int

const (
	ColorBlack Color = iota
	ColorBlue
)

func (c Color) String() string {
	switch c {
	case ColorBlack:
		return "Black"

	case ColorBlue:
		return "Blue"

	default:
		panic("Invalid color given.")
	}
}
