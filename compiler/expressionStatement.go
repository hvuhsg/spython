package compiler

import "github.com/hvuhsg/spython/ast"

func (c *context) compileExpressionStatement(expStat *ast.ExpressionStatement) error {
	err := c.compile(expStat.Expression)
	return err
}
