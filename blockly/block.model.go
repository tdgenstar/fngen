package blockly

import (
	"context"
	"fn/fngen"
	"fn/layer"
)

type Block struct {
	layer.Square
	Context context.Context
}

func Create() (*Block, error) {

	return &Block{
		Square:  layer.Square{},
		Context: context.Background(),
	}, nil
}

func (r Block) Concat() fngen.Monad {
	monad := fngen.Monad{
		Context: nil,
		Di:      nil,
		ConcatFn: fngen.ConcatFn(Block{
			Square:  r.Square,
			Context: r.Context,
		}),
	}
	return monad
}
