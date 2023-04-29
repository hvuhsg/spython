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
	token.Equal:            enum.IPredEQ,
	token.GreaterThan:      enum.IPredSGT,
	token.GreaterThenEqual: enum.IPredSGE,
	token.LessThan:         enum.IPredSLT,
	token.LessThenEqual:    enum.IPredSLE,
}

var tokenToOpFloat = map[string]enum.FPred{
	token.Equal:            enum.FPredOEQ,
	token.GreaterThan:      enum.FPredOGT,
	token.GreaterThenEqual: enum.FPredOGE,
	token.LessThan:         enum.FPredOLT,
	token.LessThenEqual:    enum.FPredOLE,
}

func (c *context) compileInfixExpression(infixExp *ast.InfixExpression) error {
	if infixExp.Operator == token.Assign {
		return c.compileAssignInfix(infixExp)
	}

	if err := c.compile(infixExp.Left); err != nil {
		return err
	}
	lreg := c.popReg()

	if err := c.compile(infixExp.Right); err != nil {
		return err
	}
	rreg := c.popReg()

	if types.IsPointer(lreg.Type()) {
		ptrTyp := lreg.Type().(*types.PointerType)
		lreg = c.NewLoad(ptrTyp.ElemType, lreg)
	}
	if types.IsPointer(rreg.Type()) {
		ptrTyp := rreg.Type().(*types.PointerType)
		rreg = c.NewLoad(ptrTyp.ElemType, rreg)
	}

	var res value.Value
	switch infixExp.Operator {
	case token.Plus:
		if lreg.Type().Equal(Int) {
			res = c.NewAdd(lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.NewFAdd(lreg, rreg)
		}
	case token.Minus:
		if lreg.Type().Equal(Int) {
			res = c.NewSub(lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.NewFSub(lreg, rreg)
		}
	case token.Asterisk:
		if types.IsInt(lreg.Type()) {
			res = c.NewMul(lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.NewFMul(lreg, rreg)
		}
	case token.Slash:
		if types.IsInt(lreg.Type()) {
			res = c.NewSDiv(lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			res = c.NewFDiv(lreg, rreg)
		}
	default:
		_, ok := tokenToOpInt[infixExp.Operator]
		if !ok {
			return newError("unsupported operator", UnsupportedError, infixExp.Token)
		}

		if types.IsInt(lreg.Type()) {
			op := tokenToOpInt[infixExp.Operator]
			res = c.NewICmp(op, lreg, rreg)
		} else if types.IsFloat(lreg.Type()) {
			op := tokenToOpFloat[infixExp.Operator]
			res = c.NewFCmp(op, lreg, rreg)
		}
	}

	c.pushReg(res)
	return nil
}

func (c *context) compileAssignInfix(assignExp *ast.InfixExpression) error {
	if err := c.compile(assignExp.Right); err != nil {
		return err
	}
	reg := c.popReg()

	identifier, ok := assignExp.Left.(*ast.Identifier)
	if !ok {
		return newError("can assign only into identifier", TypeError, assignExp.Token)
	}

	varName := identifier.TokenLiteral()

	vr := c.getVar(varName)
	if vr == nil {
		// Create new variable
		vr := c.NewAlloca(reg.Type())
		vr.SetName(varName)
		c.NewStore(reg, vr)
		c.createVar(varName, vr)
	} else {
		// Check if reg type is identical to var type
		if reg.Type().String()+"*" != vr.Type().String() {
			return newError(fmt.Sprintf("can not assign type %s into %s", reg.Type().String(), varName), TypeError, assignExp.Token)
		}
		c.NewStore(reg, vr)
	}

	return nil
}
