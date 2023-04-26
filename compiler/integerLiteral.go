package compiler

import (
	"github.com/hvuhsg/spython/ast"
	"github.com/llir/llvm/ir/constant"
)

func (c *compiler) compileIntegerLiteral(intLit *ast.IntegerLiteral) error {
	cnst := constant.NewInt(Int, intLit.Value)
	c.cstate.pushReg(cnst)
	return nil
}
