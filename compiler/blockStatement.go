package compiler

import (
	"github.com/hvuhsg/spython/ast"
)

func (c *context) compileBlockStatement(blockStat *ast.BlockStatement) error {
	for _, statement := range blockStat.Statements {
		err := c.Compile(statement)
		if err != nil {
			return err
		}
	}

	return nil
}
