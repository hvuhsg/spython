package compiler

import (
	"github.com/hvuhsg/spython/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
)

func (c *compiler) compileProgram(program *ast.Program) error {
	mainM := ir.NewModule()
	mainF := mainM.NewFunc("main", Int)
	entryB := mainF.NewBlock("")
	c.cstate.module = mainM
	c.cstate.function = mainF
	c.cstate.block = entryB
	for _, statement := range program.Statements {
		err := c.Compile(statement)
		if err != nil {
			return err
		}
	}
	c.cstate.block.NewRet(constant.NewInt(Int, 0))

	return nil
}
