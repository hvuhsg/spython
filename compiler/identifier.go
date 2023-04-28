package compiler

import (
	"fmt"

	"github.com/hvuhsg/spython/ast"
	"github.com/llir/llvm/ir/types"
)

func (c *compiler) compileIdentifier(ident *ast.Identifier) error {
	variable := c.cstate.getVariable(ident.TokenLiteral())
	if variable == nil {
		return c.newError(fmt.Sprintf("variable %s is not defined", ident.TokenLiteral()), ident.Token)
	}

	// derefrence variable value into register
	var typ types.Type
	if types.IsPointer(variable.Type()) {
		ptr := variable.Type().(*types.PointerType)
		typ = ptr.ElemType
	} else {
		typ = variable.Type()
	}

	value := c.newLoad(typ, variable)
	c.cstate.pushReg(value)
	return nil
}
