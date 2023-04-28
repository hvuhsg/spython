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

func (c *compiler) newLoad(t types.Type, v value.Value) *ir.InstLoad {
	return c.cstate.block.NewLoad(t, v)
}

func (c *compiler) newBlock(name string) *ir.Block {
	return c.cstate.function.NewBlock(name)
}

// func (c *compiler) newFunction(name string, retTyp types.Type, params ...*ir.Param) *ir.Func {
// 	return c.cstate.module.NewFunc(name, retTyp, params...)
// }
