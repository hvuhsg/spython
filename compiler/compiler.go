package compiler

import (
	"github.com/hvuhsg/spython/ast"
)

type scope struct {
	localVars     []string
	priviousScope *scope
	functions     []string
}

type compiler struct {
	globalVars   []string
	currentScope *scope
}

func New() *compiler {
	c := &compiler{}
	c.currentScope = &scope{}

	return c
}

func (c *compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, statement := range node.Statements {
			err := c.Compile(statement)
			if err != nil {
				return nil
			}
		}
	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
		// c.emit(code.OpPop)
	}

	return nil
}
