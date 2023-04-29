package compiler

import (
	"fmt"

	"github.com/hvuhsg/spython/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

func (c *context) compileFunctionLiteral(funcLit *ast.FunctionLiteral) error {
	name := funcLit.TokenLiteral()
	retTyp, ok := nameToType[funcLit.ReturnType.Value]
	if !ok {
		return newError(fmt.Sprintf("return type for function '%s' is not a valid type", name), NameError, funcLit.Token)
	}

	params := make([]*ir.Param, 0)
	for _, param := range funcLit.Parameters {
		paramName := param.TokenLiteral()
		paramTyp, ok := nameToType[param.Type.Value]
		if !ok {
			return newError(fmt.Sprintf("parameter type '%s' is not a valid type", paramName), NameError, funcLit.Token)
		}
		params = append(params, ir.NewParam(paramName, paramTyp))
	}

	fn := c.mod.NewFunc(name, retTyp, params...)
	block := fn.NewBlock("entry_" + name)
	ctx := newContext(c.mod, fn, block)

	// Load params into function local vars
	for _, param := range params {
		ctx.createVar(param.Name(), param)
	}

	if err := ctx.compile(funcLit.Body); err != nil {
		return err
	}

	return nil
}

func (c *context) compileReturnStatement(retStat *ast.ReturnStatement) error {
	if err := c.compile(retStat.ReturnValue); err != nil {
		return err
	}
	retVal := c.peekReg()

	// Check declered return type vs actual return type
	if !c.fn.Sig.RetType.Equal(retVal.Type()) {
		return newError(fmt.Sprintf("function '%s' declered return type '%s' is not matching actual return type '%s'", c.fn.Name(), c.fn.Sig.RetType.String(), retVal.Type().String()), TypeError, retStat.Token)
	}

	c.NewRet(retVal)

	return nil
}

func (c *context) compileCallExpression(callExp *ast.CallExpression) error {
	funcName := callExp.Function.TokenLiteral()
	var callee value.Value
	for _, fn := range c.mod.Funcs {
		if fn.Name() == callExp.Function.TokenLiteral() {
			callee = fn
		}
	}
	if callee == nil {
		return newError(fmt.Sprintf("function '%s' was not found", funcName), NameError, callExp.Token)
	}

	args := make([]value.Value, 0)
	for _, arg := range callExp.Arguments {
		if err := c.compile(arg); err != nil {
			return err
		}
		argReg := c.popReg()
		args = append(args, argReg)
	}

	retVal := c.NewCall(callee, args...)
	c.pushReg(retVal)

	return nil
}
