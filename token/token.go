package token

const (
	// primitives
	EOF           = iota
	Identifier    = iota
	StringLiteral = iota
	IntLiteral    = iota
	COLON         = iota
	COMA          = iota
	LEFT_PARAN    = iota
	RIGHT_PARAN   = iota

	// infix
	PLUS       = iota
	MINUS      = iota
	SLASH      = iota
	ASTRIKS    = iota
	ASSIGN     = iota
	EQUALS     = iota
	NOT_EQUALS = iota
	BIGGER     = iota
	SMALLER    = iota
	BIGGER_EQ  = iota
	SMALLER_EQ = iota

	// statements
	If     = iota
	Def    = iota
	Return = iota
)

type Token struct {
	Type  int
	Value string
	Row   int
	Col   int
	Tab   int
}

func New(typ int, val string, row int, col int, tab int) Token {
	return Token{Type: typ, Value: val, Row: row, Col: col, Tab: tab}
}

func Null() Token {
	return Token{Type: -1}
}

func IsNull(tok Token) bool {
	return tok.Type == -1
}
