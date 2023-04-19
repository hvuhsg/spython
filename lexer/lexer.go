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

	lexer.registerRegexMatcher("if", token.If)
	lexer.registerRegexMatcher("def", token.Def)
	lexer.registerRegexMatcher("return", token.Return)

	lexer.registerRegexMatcher("==", token.EQUALS)
	lexer.registerRegexMatcher("==", token.NOT_EQUALS)
	lexer.registerRegexMatcher("<=", token.SMALLER_EQ)
	lexer.registerRegexMatcher(">=", token.BIGGER_EQ)
	lexer.registerRegexMatcher(`\+`, token.PLUS)
	lexer.registerRegexMatcher("-", token.MINUS)
	lexer.registerRegexMatcher("*", token.ASTRIKS)
	lexer.registerRegexMatcher("/", token.SLASH)
	lexer.registerRegexMatcher(`=`, token.ASSIGN)
	lexer.registerRegexMatcher(":", token.COLON)
	lexer.registerRegexMatcher(",", token.COMA)
	lexer.registerRegexMatcher(`\(`, token.LEFT_PARAN)
	lexer.registerRegexMatcher(`\)`, token.RIGHT_PARAN)

	return lexer
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	if l.isEnd() {
		return l.newToken(token.EOF, "EOF")
	}

	currentData := l.currentData()

	for _, matcher := range l.matchers {
		tok := matcher(currentData, l)
		if token.IsNull(tok) {
			continue
		}

		l.cursor += len(tok.Value)
		l.col += len(tok.Value)

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
		case '\n':
			l.cursor += 1
			l.col = 0
			l.row += 1
			l.tab = 0
		case '\t':
			l.cursor += 1
			l.col += 1
			l.tab += 0
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

func (l *Lexer) registerRegexMatcher(pattern string, tokenTyp int) {
	re := regexp.MustCompile(`^` + pattern)
	matcher := func(data string, l *Lexer) token.Token {
		res := re.FindString(data)
		if res == "" {
			return token.Null()
		}

		return l.newToken(tokenTyp, res)
	}

	l.matchers = append(l.matchers, matcher)
}

func (l *Lexer) newToken(typ int, val string) token.Token {
	return token.New(typ, val, l.row, l.col, l.tab)
}
