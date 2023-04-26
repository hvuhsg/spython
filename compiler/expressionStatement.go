package compiler

import "github.com/hvuhsg/spython/ast"

func (c *compiler) compileExpressionStatement(expStat *ast.ExpressionStatement) error {
	err := c.Compile(expStat.Expression)
	return err
}
