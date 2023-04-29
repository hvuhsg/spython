package compiler

import (
	"github.com/hvuhsg/spython/ast"
	"github.com/llir/llvm/ir/constant"
)

func (c *context) compileProgram(program *ast.Program) error {
	for _, statement := range program.Statements {
		err := c.compile(statement)
		if err != nil {
			return err
		}
	}

	retVal := c.popReg()
	if retVal != nil {
		c.NewRet(retVal)
	} else {
		c.NewRet(constant.NewInt(Int, 0))
	}

	return nil
}
