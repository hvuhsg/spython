package compiler

import (
	"fmt"

	"github.com/hvuhsg/spython/ast"
)

func (c *compiler) compileIdentifier(ident *ast.Identifier) error {
	variable := c.cstate.GetVariable(ident.TokenLiteral())
	if variable == nil {
		return c.newError(fmt.Sprintf("variable %s is not defined", ident.TokenLiteral()), ident.Token)
	}

	c.cstate.pushReg(variable)
	return nil
}
