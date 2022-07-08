package layer

type Line struct {
	name  string
	Start Point
	End   Point
}

func (r Line) LengthToPoint() (float64, error) {
	return 0, err("not impl")
}

func (r Line) Name() (string, error) {
	if r.name == "" {
		return "", err("not name")
	}
	return r.name, nil
}
