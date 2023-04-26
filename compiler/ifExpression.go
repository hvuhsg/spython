package compiler

import (
	"github.com/hvuhsg/spython/ast"
	"github.com/llir/llvm/ir"
)

func (c *compiler) compileIfExpression(ifExp *ast.IfExpression) error {
	end := c.cstate.function.NewBlock("")

	// compile condition
	if err := c.Compile(ifExp.Condition); err != nil {
		return err
	}
	cond := c.cstate.popReg()

	if err := c.Compile(ifExp.Consequence); err != nil {
		return err
	}
	target := c.cstate.popBlock()
	target.NewBr(end)

	// create else brach
	var ault *ir.Block = nil
	if ifExp.Alternative != nil {
		if err := c.Compile(ifExp.Alternative); err != nil {
			return err
		}
		ault = c.cstate.popBlock()
		ault.NewBr(end)
	}

	// create branch
	c.cstate.block.NewCondBr(cond, target, ault)

	c.cstate.block = end
	return nil
}
