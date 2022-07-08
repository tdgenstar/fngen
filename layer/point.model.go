package layer

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func (r Point) String() (string, error) {
	err1 := r.validate()
	if err1 != nil {
		return "", err1
	}

	return fmt.Sprintf("%f,%f", r.X, r.Y), nil
}

func (r Point) Draw() ([]Point, error) {
	err1 := r.validate()
	if err1 != nil {
		return []Point{}, err1
	}

	return []Point{r}, nil
}

func (r Point) validate() error {
	if r.X == 0 || r.Y == 0 {
		return err("0,0 is nil")
	}
	return nil
}
