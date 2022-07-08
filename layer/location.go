package layer

import (
	"fmt"
)

type Location struct {
	X float64
	Y float64
}

func (r Location) String() (string, error) {
	if r.X == 0 || r.Y == 0 {
		return "", err("0,0 is nil")
	}
	return fmt.Sprintf("%f,%f", r.X, r.Y), nil
}
