package layer

type Line struct {
	name  string
	Start Point
	End   Point
}

type Square struct {
	name  string
	Start Point
	End   Point
}

func (r Line) LengthToPoint() (float64, error) {
	return 0, err("not impl")
}

func (r Line) Name() (string, error) {
	if r.name == "" {
		return "", err("empty name")
	}
	return r.name, nil
}

func (r Line) Map() []Line {
	if r.name == "" {
		return []Line{}
	}
	return []Line{r}
}

func (r Square) Name() (string, error) {
	if r.name == "" {
		return "", err("empty name")
	}
	return r.name, nil
}
