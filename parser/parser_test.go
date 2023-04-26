package parser

import (
	"testing"

	"github.com/hvuhsg/spython/lexer"
)

func TestFibonaci(t *testing.T) {
	lexer := lexer.New("def fib(n: int = 5):\n\tif n == 0 or n == 1:\n\t\treturn 1\n\treturn fib(n-1) + fib(n-2)")
	parser := New(&lexer)
	program := parser.ParseProgram()

	if program.String() != "def fib(n: int = 5):\n\tif ((n == 0) or (n == 1)):\n\t\treturn 1\n\treturn (fib((n - 1)) + fib((n - 2)))\n" {
		t.Error("Finbonaci program does not parsed correctly")
	}
}

func TestWhileLoop(t *testing.T) {
	lexer := lexer.New("while i > 6:\n\tprint(i)")
	parser := New(&lexer)
	program := parser.ParseProgram()

	if program.String() != "while (i > 6):\n\tprint(i)" {
		t.Error("While loop was not parsed correctly")
	}
}