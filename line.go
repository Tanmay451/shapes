package shapes

import "fmt"

type Line struct {
	Length float64
}

func (l Line) String() string {
	return fmt.Sprintf("Line(%0.2f)", l.Length)
}

func (l Line) Area() float64 {
	return 0
}
