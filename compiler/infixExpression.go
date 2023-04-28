package compiler

import (
	"fmt"

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
		ptrTyp := lreg.Type().(*types.PointerType)
		lreg = c.newLoad(ptrTyp.ElemType, lreg)
	}
	if types.IsPointer(rreg.Type()) {
		ptrTyp := rreg.Type().(*types.PointerType)
		rreg = c.newLoad(ptrTyp.ElemType, rreg)
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

	varName := identifier.TokenLiteral()

	vr := c.cstate.getVariable(varName)
	if vr == nil {
		// Create new variable
		vr := c.cstate.block.NewAlloca(reg.Type())
		vr.SetName(varName)
		c.cstate.block.NewStore(reg, vr)
		c.cstate.addVariable(varName, vr)
	} else {
		// Check if reg type is identical to var type
		if reg.Type().String()+"*" != vr.Type().String() {
			return c.newError(fmt.Sprintf("can not assign type %s into %s", reg.Type().String(), varName), assignExp.Token)
		}
		c.cstate.block.NewStore(reg, vr)
	}

	return nil
}
