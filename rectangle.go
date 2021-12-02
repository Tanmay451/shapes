package shapes

import "fmt"

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle Width: (%0.2f), Length: (%0.2f), ", r.Width, r.Length)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}
