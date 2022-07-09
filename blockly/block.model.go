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

func Create(ctx context.Context) (*Block, error) {

	return &Block{
		Square:  layer.Square{},
		Context: ctx,
	}, nil
}

func (r Block) Concat() fngen.Monad {
	monad := fngen.Monad{
		Context: r.Context,
		Di:      nil,
		ConcatFn: fngen.ConcatFn(Block{
			Square:  r.Square,
			Context: r.Context,
		}),
	}
	return monad
}
