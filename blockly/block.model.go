package blockly

import (
	"fn/fngen"
	"fn/layer"
)

type Block struct {
	layer.Square
}

func Create() (*Block, error) {
	return nil, nil
}

func (r Block) Concat() fngen.Monad {
	return fngen.Monad{}
}
