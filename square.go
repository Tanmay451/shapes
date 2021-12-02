package shapes

import "fmt"

type Square struct {
	Side float64
}

func (s Square) String() string {
	return fmt.Sprintf("Square(%0.2f)", s.Side)
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}
