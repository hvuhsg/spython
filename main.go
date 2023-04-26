package main

import (
	"fmt"
	"os"

	"github.com/hvuhsg/spython/compiler"
	"github.com/hvuhsg/spython/lexer"
	"github.com/hvuhsg/spython/parser"
)

func main() {
	code := "a = 5\nif a == 5:\n\ta = 0\nelse:\n\ta = 1\n"
	fmt.Println("Code:")
	fmt.Println(code)
	fmt.Println()

	lexer := lexer.New(code)
	parser := parser.New(&lexer)
	compiler := compiler.New()

	ast := parser.ParseProgram()
	fmt.Println("Parser:")
	fmt.Println(ast.String())
	fmt.Println()

	err := compiler.Compile(ast)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	llvmCode := compiler.IR()
	fmt.Println("LLVM IR:")
	fmt.Println(llvmCode)

	os.WriteFile("./code.ll", []byte(llvmCode), 0644)
}
