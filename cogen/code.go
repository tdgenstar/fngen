package cogen

type Cogen struct {
	Group
	Name  string
	Order int
	Code  string
}

type Group int64
