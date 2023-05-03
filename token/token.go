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
	Float      = "Float"      // 3.14
	String     = "String"     // "x", "y"
	None       = "None"       // None

	// Operators
	Assign   = "="
	Mod      = "%"
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
	Arrow     = "->"

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
