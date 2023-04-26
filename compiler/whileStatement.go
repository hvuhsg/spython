package compiler

import "github.com/hvuhsg/spython/ast"

func (c *compiler) compileWhileExpression(whileExp *ast.WhileExpression) error {
	end := c.cstate.function.NewBlock("")
	priv := c.cstate.block

	// Create while block
	while := c.cstate.function.NewBlock("")
	c.cstate.block = while
	if err := c.Compile(whileExp.Condition); err != nil {
		return err
	}
	cond := c.cstate.popReg()

	// Create loop block
	c.cstate.block = priv
	if err := c.Compile(whileExp.Consequence); err != nil {
		return err
	}
	loop := c.cstate.popBlock()
	loop.NewBr(while)

	// Create loop condition
	while.NewCondBr(cond, loop, end)

	// Jump to while
	priv.NewBr(while)

	c.cstate.block = end
	return nil
}
