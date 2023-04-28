package compiler

import "github.com/hvuhsg/spython/ast"

func (c *compiler) compileBlockStatement(blockStat *ast.BlockStatement) error {
	priv := c.cstate.block

	b := c.newBlock("")
	c.cstate.block = b

	for _, statement := range blockStat.Statements {
		err := c.Compile(statement)
		if err != nil {
			return err
		}
	}

	c.cstate.pushBlock(b)

	c.cstate.block = priv
	return nil
}
