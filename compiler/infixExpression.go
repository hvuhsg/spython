package compiler

import (
	"github.com/hvuhsg/spython/ast"
	"github.com/hvuhsg/spython/token"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var tokenToOpInt = map[string]enum.IPred{
	token.Assign:           enum.IPredEQ,
	token.GreaterThan:      enum.IPredSGT,
	token.GreaterThenEqual: enum.IPredSGE,
	token.LessThan:         enum.IPredSLT,
	token.LessThenEqual:    enum.IPredSLE,
}

var tokenToOpFloat = map[string]enum.FPred{
	token.Assign:           enum.FPredOEQ,
	token.GreaterThan:      enum.FPredOGT,
	token.GreaterThenEqual: enum.FPredOGE,
	token.LessThan:         enum.FPredOLT,
	token.LessThenEqual:    enum.FPredOLE,
}

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

	if types.IsPointer(lreg.Type()) {
		if lreg.Type().String() == "i64*" {
			lreg = c.cstate.block.NewLoad(Int, lreg)
			// lreg = c.cstate.block.NewPtrToInt(lreg, Int)
		} else if lreg.Type().String() == "float*" {
			lreg = c.cstate.block.NewLoad(Float, lreg)
		}
	}
	if types.IsPointer(rreg.Type()) {
		if rreg.Type().String() == "i64*" {
			rreg = c.cstate.block.NewLoad(Int, lreg)
			// rreg = c.cstate.block.NewPtrToInt(rreg, Int)
		} else if rreg.Type().String() == "float*" {
			rreg = c.cstate.block.NewLoad(Float, rreg)
		}
	}

	var res value.Value
	switch infixExp.Operator {
	case token.Plus:
		if lreg.Type().Equal(Int) {
			res = c.cstate.block.NewAdd(lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.cstate.block.NewFAdd(lreg, rreg)
		}
	case token.Minus:
		if lreg.Type().Equal(Int) {
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
	default:
		_, ok := tokenToOpInt[infixExp.Operator]
		if !ok {
			return c.newError("unsupported operator", infixExp.Token)
		}

		if types.IsInt(lreg.Type()) {
			op := tokenToOpInt[infixExp.Operator]
			res = c.cstate.block.NewICmp(op, lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			op := tokenToOpFloat[infixExp.Operator]
			res = c.cstate.block.NewFCmp(op, lreg, rreg)
		}
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
