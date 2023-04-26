package compiler

import (
	"fmt"

	"github.com/hvuhsg/spython/ast"
	"github.com/hvuhsg/spython/token"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var Int = types.I64
var Float = types.Float

type state struct {
	variables  map[string]value.Value
	token      token.Token
	module     *ir.Module
	function   *ir.Func
	block      *ir.Block
	regStack   []value.Value
	blockStack []*ir.Block
}

func newState() *state {
	s := &state{}
	s.regStack = make([]value.Value, 0)
	s.blockStack = make([]*ir.Block, 0)
	s.variables = make(map[string]value.Value)

	return s
}

func (s *state) AddVariable(name string, val value.Value) {
	s.variables[name] = val
}

func (s *state) GetVariable(name string) value.Value {
	return s.variables[name]
}

func (s *state) pushBlock(block *ir.Block) {
	s.blockStack = append(s.blockStack, block)
}

func (s *state) popBlock() *ir.Block {
	if len(s.blockStack) == 0 {
		return nil
	}

	index := len(s.blockStack) - 1
	block := (s.blockStack)[index]
	s.blockStack = (s.blockStack)[:index]
	return block
}

func (s *state) pushReg(val value.Value) {
	s.regStack = append(s.regStack, val)
}

func (s *state) popReg() value.Value {
	if len(s.regStack) == 0 {
		return nil
	}

	index := len(s.regStack) - 1
	reg := (s.regStack)[index]
	s.regStack = (s.regStack)[:index]
	return reg
}

type compiler struct {
	cstate *state
}

func New() *compiler {
	c := &compiler{}
	c.cstate = newState()

	return c
}

func (c *compiler) IR() string {
	return fmt.Sprintln(c.cstate.module)
}

func (c *compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		if err := c.compileProgram(node); err != nil {
			return err
		}
	case *ast.ExpressionStatement:
		if err := c.compileExpressionStatement(node); err != nil {
			return err
		}
	case *ast.InfixExpression:
		if err := c.compileInfixExpression(node); err != nil {
			return err
		}
	case *ast.IntegerLiteral:
		if err := c.compileIntegerLiteral(node); err != nil {
			return err
		}
	case *ast.FloatLiteral:
		if err := c.compileFloatLiteral(node); err != nil {
			return err
		}
	case *ast.Identifier:
		if err := c.compileIdentifier(node); err != nil {
			return err
		}
	case *ast.IfExpression:
		if err := c.compileIfExpression(node); err != nil {
			return err
		}
	case *ast.WhileExpression:
		if err := c.compileWhileExpression(node); err != nil {
			return err
		}
	case *ast.BlockStatement:
		if err := c.compileBlockStatement(node); err != nil {
			return err
		}
	default:
		fmt.Println(node)
		return fmt.Errorf("node not supported")
	}

	return nil
}

func (c *compiler) newError(s string, tok token.Token) error {
	c.cstate.token = tok
	return fmt.Errorf(s)
}
