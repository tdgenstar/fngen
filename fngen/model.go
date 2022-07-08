package fngen

import "context"

type Monad struct {
	key     []string
	Context context.Context
	Di      []Monad
}

type FuncBase interface {
	Map() Monad
	Reduce() Monad

	Collect() Monad
}

func (m Monad) Map() Monad {
	return m
}

func (m Monad) Reduce() Monad {
	return m
}

func (m Monad) Collect() Monad {
	var keys []string
	for _, v := range m.Di {
		keys = append(keys, v.key...)
	}
	var x = Monad{
		key:     keys,
		Context: m.Context,
		Di:      nil,
	}
	return x
}
