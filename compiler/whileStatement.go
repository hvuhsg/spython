package compiler

import (
	"github.com/hvuhsg/spython/ast"
)

func (c *context) compileWhileExpression(whileExp *ast.WhileExpression) error {
	endwhile := c.newContext("while.end")

	// Create while block
	condition := c.newContext("while.condition")
	if err := condition.Compile(whileExp.Condition); err != nil {
		return err
	}
	cond := condition.popReg()

	// Create loop block
	loop := c.newContext("while.loop")
	if err := loop.Compile(whileExp.Consequence); err != nil {
		return err
	}
	loop.NewBr(endwhile.Block)

	// Create loop condition
	condition.NewCondBr(cond, loop.Block, endwhile.Block)

	// Jump to while
	c.NewBr(condition.Block)

	// Continue with endif block
	c.Block = endwhile.Block
	return nil
}
