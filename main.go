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
	code := "def fib(n: int) -> int:\n\ta = 0\n\tb = 1\n\twhile n > 0:\n\t\tn = n - 1\n\t\tb = a + b\n\t\ta = b - a\n\n\treturn b\nreturn fib(40)"

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
