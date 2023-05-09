package main

import (
	"fmt"
	"os"

	"github.com/hvuhsg/spython/compiler"
	"github.com/hvuhsg/spython/lexer"
	"github.com/hvuhsg/spython/parser"
)

func main() {
	// Read SPython code
	// if len(os.Args) < 2 {
	// 	fmt.Println("Please specify a file path as an argument.")
	// 	return
	// }

	// path := os.Args[1]

	// data, errr := os.ReadFile(path)
	// if errr != nil {
	// 	fmt.Println("Error reading file:", errr)
	// 	return
	// }

	// code := string(data)
	code := "def b():\n\tif 1 > 2:\n\t\ta = 1\nelse:\n\t\ta = 2\n\n\treturn 2"

	fmt.Println("Code:")
	fmt.Println(code)
	fmt.Println()

	lexer := lexer.New(code)
	parser := parser.New(&lexer)
	compiler := compiler.New()

	ast := parser.ParseProgram()

	if len(parser.Errors()) > 0 {
		fmt.Printf("Parser errors: %v\n", parser.Errors())
		os.Exit(1)
	}

	fmt.Println("Parser:")
	fmt.Println(ast.String())
	fmt.Println()

	err := compiler.Compile(ast)
	if err != nil {
		fmt.Printf("Compiler error: %s\n", err.Error())
		os.Exit(1)
	}

	llvmCode := compiler.IR()
	fmt.Println("LLVM IR:")
	fmt.Println(llvmCode)

	os.WriteFile("./code.ll", []byte(llvmCode), 0644)
}
