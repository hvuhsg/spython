package parser

import (
	"github.com/hvuhsg/spython/token"

	"github.com/hvuhsg/spython/lexer"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New() Parser {
	return Parser{}
}
