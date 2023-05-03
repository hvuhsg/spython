package lexer

import (
	"regexp"

	"github.com/hvuhsg/spython/token"
)

type Lexer struct {
	data   string
	cursor int
	row    int
	col    int
	tab    int

	matchers []func(data string, l *Lexer) token.Token
}

func New(data string) Lexer {
	lexer := Lexer{data: data}

	lexer.registerSimpleMatcher("None", token.None)
	lexer.registerSimpleMatcher("if", token.If)
	lexer.registerSimpleMatcher("else", token.Else)
	lexer.registerSimpleMatcher("def", token.Function)
	lexer.registerSimpleMatcher("return", token.Return)
	lexer.registerSimpleMatcher("while", token.While)
	lexer.registerRegexMatcher(`[0-9]*\.[0-9]+`, token.Float)
	lexer.registerRegexMatcher(`\d*`, token.Int)
	lexer.registerSimpleMatcher("\n", token.ENDL)

	// logic gates
	lexer.registerSimpleMatcher("or", token.Or)
	lexer.registerSimpleMatcher("and", token.And)

	lexer.registerSimpleMatcher("%", token.Mod)
	lexer.registerSimpleMatcher("->", token.Arrow)
	lexer.registerSimpleMatcher("==", token.Equal)
	lexer.registerSimpleMatcher("!=", token.NotEqual)
	lexer.registerSimpleMatcher(">=", token.GreaterThenEqual)
	lexer.registerSimpleMatcher("<=", token.LessThenEqual)
	lexer.registerSimpleMatcher("<", token.LessThan)
	lexer.registerSimpleMatcher(">", token.GreaterThan)
	lexer.registerSimpleMatcher("+", token.Plus)
	lexer.registerSimpleMatcher("-", token.Minus)
	lexer.registerSimpleMatcher("*", token.Asterisk)
	lexer.registerSimpleMatcher("/", token.Slash)
	lexer.registerSimpleMatcher(`=`, token.Assign)
	lexer.registerSimpleMatcher(":", token.Colon)
	lexer.registerSimpleMatcher(",", token.Comma)
	lexer.registerSimpleMatcher("(", token.LeftParen)
	lexer.registerSimpleMatcher(")", token.RightParen)

	lexer.registerRegexMatcher("[a-zA-Z]([a-zA-Z0-9]*)", token.Identifier)

	return lexer
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	if l.isEnd() {
		return l.newToken(token.EOF, "")
	}

	currentData := l.currentData()

	for _, matcher := range l.matchers {
		tok := matcher(currentData, l)
		if tok.Type == token.Illegal {
			continue
		}

		l.cursor += len(tok.Literal)
		l.col += len(tok.Literal)

		if tok.Type == token.ENDL {
			l.col = 0
			l.row += 1
			l.tab = 0
			tok.Col = l.col
			tok.Row += l.row
			tok.Tab = l.tab
		}

		return tok
	}

	panic("Got: " + l.currentData())
}

// Check if current data is whitespace and skip it until there are't any more
func (l *Lexer) skipWhitespace() {
	hasWhitespace := true

	for hasWhitespace && !l.isEnd() {
		switch l.currentData()[0] {
		case ' ':
			l.cursor += 1
			l.col += 1
		case '\t':
			l.cursor += 1
			l.col += 1
			l.tab += 1
		default:
			hasWhitespace = false
		}
	}
}

func (l *Lexer) isEnd() bool {
	return len(l.data) == l.cursor
}

func (l *Lexer) currentData() string {
	return l.data[l.cursor:]
}

func (l *Lexer) registerRegexMatcher(pattern string, tokenTyp token.TokenType) {
	re := regexp.MustCompile(`^` + pattern)
	matcher := func(data string, l *Lexer) token.Token {
		res := re.FindString(data)
		if res == "" {
			return l.newToken(token.Illegal, "")
		}

		return l.newToken(tokenTyp, res)
	}

	l.matchers = append(l.matchers, matcher)
}

func (l *Lexer) newToken(typ token.TokenType, val string) token.Token {
	return token.Token{Type: typ, Literal: val, Row: l.row, Col: l.col, Tab: l.tab}
}

func (l *Lexer) registerSimpleMatcher(word string, tokenTyp token.TokenType) {
	matcher := func(data string, l *Lexer) token.Token {
		if len(data) < len(word) || data[:len(word)] != word {
			return l.newToken(token.Illegal, "")
		}

		return l.newToken(tokenTyp, word)
	}

	l.matchers = append(l.matchers, matcher)
}
