package layer

import (
	"fmt"
)

type Location struct {
	X float64
	Y float64
}

func (r Location) String() (string, error) {
	err1 := r.validate()
	if err1 != nil {
		return "", err1
	}

	return fmt.Sprintf("%f,%f", r.X, r.Y), nil
}

func (r Location) Draw() ([]Location, error) {
	err1 := r.validate()
	if err1 != nil {
		return nil, err1
	}

	return []Location{r}, nil
}

func (r Location) validate() error {
	if r.X == 0 || r.Y == 0 {
		return err("0,0 is nil")
	}
	return nil
}
