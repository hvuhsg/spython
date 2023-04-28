package compiler

import (
	"testing"

	"github.com/hvuhsg/spython/lexer"
	"github.com/hvuhsg/spython/parser"
)

func TestVariableNotDefined(t *testing.T) {
	l := lexer.New("a = b")
	p := parser.New(&l)
	c := New()
	ast := p.ParseProgram()

	if len(p.Errors()) != 0 {
		t.Errorf("Got parsing errors %v", p.Errors())
	}

	err := c.Compile(ast)
	if err == nil || err.Error() != "variable b is not defined" {
		t.Errorf("Expecting an undefined error for variable 'b' got %s", err.Error())
	}
}

func TestWrongTypeAssign(t *testing.T) {
	l := lexer.New("a = 1\na=0.6")
	p := parser.New(&l)
	c := New()
	ast := p.ParseProgram()

	if len(p.Errors()) != 0 {
		t.Errorf("Got parsing errors %v", p.Errors())
	}

	err := c.Compile(ast)
	if err == nil || err.Error() != "can not assign type float into a" {
		t.Errorf("Expecting a wrong type assign error got %s", err.Error())
	}
}
