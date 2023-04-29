package compiler

import (
	"github.com/hvuhsg/spython/ast"
)

func (c *context) compileIfExpression(ifExp *ast.IfExpression) error {
	endif := c.newContext("endif")

	// compile condition
	if err := c.compile(ifExp.Condition); err != nil {
		return err
	}
	cond := c.popReg()

	ifCtx := c.newContext("if.then")
	if err := ifCtx.compile(ifExp.Consequence); err != nil {
		return err
	}

	// only jump to endif if there is no return
	if reVal := ifCtx.popReg(); reVal == nil {
		ifCtx.NewBr(endif.Block)
	}

	// create else brach
	var elseCtx *context = endif
	if ifExp.Alternative != nil {
		elseCtx = c.newContext("if.else")
		if err := elseCtx.compile(ifExp.Alternative); err != nil {
			return err
		}
		elseCtx.NewBr(endif.Block)
	}

	// create branch
	c.NewCondBr(cond, ifCtx.Block, elseCtx.Block)

	// Continue with endif block
	c.Block = endif.Block
	return nil
}
