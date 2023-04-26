package compiler

import (
	"github.com/hvuhsg/spython/ast"
	"github.com/hvuhsg/spython/token"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *compiler) compileInfixExpression(infixExp *ast.InfixExpression) error {
	if infixExp.Operator == token.Assign {
		return c.compileAssignInfix(infixExp)
	}

	if err := c.Compile(infixExp.Left); err != nil {
		return err
	}
	lreg := c.cstate.popReg()

	if err := c.Compile(infixExp.Right); err != nil {
		return err
	}

	rreg := c.cstate.popReg()

	var res value.Value
	switch infixExp.Operator {
	case token.Plus:
		if types.IsInt(lreg.Type()) {
			res = c.cstate.block.NewAdd(lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.cstate.block.NewFAdd(lreg, rreg)
		}
	case token.Minus:
		if types.IsInt(lreg.Type()) {
			res = c.cstate.block.NewSub(lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.cstate.block.NewFSub(lreg, rreg)
		}
	case token.Asterisk:
		if types.IsInt(lreg.Type()) {
			res = c.cstate.block.NewMul(lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.cstate.block.NewFMul(lreg, rreg)
		}
	case token.Slash:
		if types.IsInt(lreg.Type()) {
			res = c.cstate.block.NewSDiv(lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.cstate.block.NewFDiv(lreg, rreg)
		}
	case token.Equal:
		if types.IsPointer(lreg.Type()) {
			if types.IsInt(lreg.Type()) {
				lreg = c.cstate.block.NewPtrToInt(lreg, Int)
			} else {
				lreg = c.cstate.block.NewLoad(Float, lreg)
			}
		}
		if types.IsPointer(rreg.Type()) {
			if types.IsInt(rreg.Type()) {
				rreg = c.cstate.block.NewPtrToInt(rreg, Int)
			} else {
				rreg = c.cstate.block.NewLoad(Float, rreg)
			}
		}
		if types.IsInt(lreg.Type()) {
			res = c.cstate.block.NewICmp(enum.IPredEQ, lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.cstate.block.NewFCmp(enum.FPredOEQ, lreg, rreg)
		}

	default:
		return c.newError("operator not supported", infixExp.Token)
	}

	c.cstate.pushReg(res)
	return nil
}

func (c *compiler) compileAssignInfix(assignExp *ast.InfixExpression) error {
	if err := c.Compile(assignExp.Right); err != nil {
		return err
	}
	reg := c.cstate.popReg()

	identifier, ok := assignExp.Left.(*ast.Identifier)
	if !ok {
		return c.newError("can assign only into identifier", assignExp.Token)
	}

	vr := c.cstate.GetVariable(identifier.TokenLiteral())
	if vr == nil {
		vr := c.cstate.block.NewAlloca(reg.Type())
		vr.SetName(identifier.TokenLiteral())
		c.cstate.block.NewStore(reg, vr)
		c.cstate.AddVariable(identifier.TokenLiteral(), vr)
	} else {
		c.cstate.block.NewStore(reg, vr)
	}

	return nil
}
