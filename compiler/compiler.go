package compiler

import (
	"fmt"

	"github.com/hvuhsg/spython/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

var Int = types.I64
var Float = types.Float
var None = types.Void

var nameToType map[string]types.Type = map[string]types.Type{
	"int":   Int,
	"float": Float,
	"None":  None,
}

type compiler struct {
	module   *ir.Module
	function *ir.Func
	ctx      *context
}

func New() *compiler {
	c := &compiler{}

	mainModule := ir.NewModule()
	c.module = mainModule

	mainFunction := mainModule.NewFunc("main", Int)
	c.function = mainFunction

	startBlock := mainFunction.NewBlock("prog_entry")
	c.ctx = newContext(mainModule, mainFunction, startBlock)

	return c
}

func (c *compiler) IR() string {
	return fmt.Sprintln(c.module)
}

func (c *compiler) Compile(prog *ast.Program) error {
	return c.ctx.compile(prog)
}

func (c *context) compile(node ast.Node) error {
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
	case *ast.FunctionLiteral:
		if err := c.compileFunctionLiteral(node); err != nil {
			return err
		}
	case *ast.ReturnStatement:
		if err := c.compileReturnStatement(node); err != nil {
			return err
		}
	case *ast.CallExpression:
		if err := c.compileCallExpression(node); err != nil {
			return err
		}
	default:
		fmt.Println(node)
		return fmt.Errorf("node not supported")
	}

	return nil
}
