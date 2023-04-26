package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Row     int
	Col     int
	Tab     int
}

const (
	Illegal = "Illegal"
	EOF     = "EOF"
	ENDL    = "\n"

	// Identifiers + Literals
	Identifier = "Identifier" // add, x ,y, ...
	Int        = "Int"        // 123456
	String     = "String"     // "x", "y"

	// Operators
	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Bang     = "!"
	Asterisk = "*"
	Slash    = "/"
	Equal    = "=="
	NotEqual = "!="
	Or       = "||"
	And      = "&&"

	LessThan         = "<"
	LessThenEqual    = "<="
	GreaterThan      = ">"
	GreaterThenEqual = ">="

	// Delimiters
	Comma     = ","
	Semicolon = ";"
	Colon     = ":"

	LeftParen    = "("
	RightParen   = ")"
	LeftBrace    = "{"
	RightBrace   = "}"
	LeftBracket  = "["
	RightBracket = "]"

	// Keywords
	Function = "Function"
	True     = "True"
	False    = "False"
	If       = "If"
	Else     = "Else"
	Return   = "Return"
	For      = "For"
	While    = "while"
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookupIdentifierType(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return Identifier
}