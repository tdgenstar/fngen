package cogen

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

type Cogen struct {
	GroupId
	Name  string
	Order int
	Code  string
	file  *File
	*Statement
}

type GroupId int64

func ExampleCode() {
	file := NewFile("fngen_gencode")
	statement := file.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello, world")),
	)
	cogen := Cogen{
		GroupId:   1,
		Name:      "a1",
		Order:     0,
		Code:      "test code",
		file:      NewFile("fngen_gencode"),
		Statement: statement,
	}

	fmt.Printf("%#v", cogen.Statement)
}
