package compiler

import (
	"github.com/hvuhsg/spython/ast"
	"github.com/llir/llvm/ir/constant"
)

func (c *context) compileIntegerLiteral(intLit *ast.IntegerLiteral) error {
	cnst := constant.NewInt(Int, intLit.Value)
	c.pushReg(cnst)
	return nil
}

func (c *context) compileFloatLiteral(floatLit *ast.FloatLiteral) error {
	cnst := constant.NewFloat(Float, floatLit.Value)
	c.pushReg(cnst)
	return nil
}
