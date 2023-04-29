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

	// data, err := os.ReadFile(path)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// code := string(data)
	code := "def fib(n: int) -> int:\n\tif n < 1:\n\t\treturn 1\n\treturn fib(n-1) + fib(n-2)\na = fib(10)"

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
