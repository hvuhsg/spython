package compiler

import "github.com/hvuhsg/spython/token"

const (
	NameError = iota
	TypeError
	UnsupportedError
)

type CompilationError int

type compileError struct {
	Msg   string
	Type  CompilationError
	Token token.Token
}

func (c compileError) Error() string {
	return c.Msg
}

func newError(msg string, typ CompilationError, tok token.Token) compileError {
	return compileError{Msg: msg, Type: typ, Token: tok}
}
