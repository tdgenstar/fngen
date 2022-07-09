package io

import (
	"context"
	"fmt"
	"fn/blockly"
	"log"
)

func Start() {
	ctx := context.Background()
	var block1, err1 = blockly.Create(ctx)
	if err1 != nil {
		log.Panicf("%v\n", err1)
	}

	fmt.Println(block1.Name())
}
