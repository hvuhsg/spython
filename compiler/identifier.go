package compiler

import (
	"fmt"

	"github.com/hvuhsg/spython/ast"
	"github.com/llir/llvm/ir/types"
)

func (c *context) compileIdentifier(ident *ast.Identifier) error {
	variable := c.getVar(ident.TokenLiteral())
	if variable == nil {
		return newError(fmt.Sprintf("variable %s is not defined", ident.TokenLiteral()), NameError, ident.Token)
	}

	// derefrence variable value into register
	if types.IsPointer(variable.Type()) {
		ptr := variable.Type().(*types.PointerType)
		typ := ptr.ElemType
		value := c.NewLoad(typ, variable)
		c.pushReg(value)
	} else {
		c.pushReg(variable)
	}

	return nil
}
