package layer

type Line struct {
	Start Point
	End   Point
}

func (r Line) LengthToPoint() (float64, error) {
	return 0, err("not impl")
}
