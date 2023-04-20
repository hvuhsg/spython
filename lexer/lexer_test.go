package lexer

import (
	"testing"

	"github.com/hvuhsg/spython/token"
)

func TestLexer(t *testing.T) {
	lexer := New("+ = \n \t(==)")

	expectedTokens := []token.Token{
		{Type: token.Plus, Literal: "+", Row: 0, Col: 0, Tab: 0},
		{Type: token.Assign, Literal: "=", Row: 0, Col: 2, Tab: 0},
		{Type: token.ENDL, Literal: "\n", Row: 1, Col: 0, Tab: 0},
		{Type: token.LeftParen, Literal: "(", Row: 1, Col: 2, Tab: 1},
		{Type: token.Equal, Literal: "==", Row: 1, Col: 3, Tab: 1},
		{Type: token.RightParen, Literal: ")", Row: 1, Col: 5, Tab: 1},
		{Type: token.EOF, Literal: "", Row: 1, Col: 6, Tab: 1},
	}

	for _, et := range expectedTokens {
		token := lexer.NextToken()

		if token.Type != et.Type {
			t.Fatalf("Expected token type %s got %s", et.Type, token.Type)
		}

		if token.Literal != et.Literal {
			t.Fatalf("Expected token value '%s' got '%s'", et.Literal, token.Literal)
		}

		if token.Row != et.Row {
			t.Fatalf("Expected token row %d got %d", et.Row, token.Row)
		}

		if token.Col != et.Col {
			t.Fatalf("Expected token col %d got %d", et.Col, token.Col)
		}

		if token.Tab != et.Tab {
			t.Fatalf("Expected token tab %d got %d", et.Tab, token.Tab)
		}
	}
}

func TestLexerIf(t *testing.T) {
	lexer := New("if a == 5:\n\ta = 4")

	expectedTokens := []token.Token{
		{Type: token.If, Literal: "if", Tab: 0},
		{Type: token.Identifier, Literal: "a", Tab: 0},
		{Type: token.Equal, Literal: "==", Tab: 0},
		{Type: token.Int, Literal: "5", Tab: 0},
		{Type: token.Colon, Literal: ":", Tab: 0},
		{Type: token.ENDL, Literal: "\n", Tab: 0},
		{Type: token.Identifier, Literal: "a", Tab: 1},
		{Type: token.Assign, Literal: "=", Tab: 1},
		{Type: token.Int, Literal: "4", Tab: 1},
	}

	for _, et := range expectedTokens {
		token := lexer.NextToken()

		if token.Type != et.Type {
			t.Fatalf("Expected token type %s got %s", et.Type, token.Type)
		}

		if token.Literal != et.Literal {
			t.Fatalf("Expected token value '%s' got '%s'", et.Literal, token.Literal)
		}

		if token.Tab != et.Tab {
			t.Fatalf("Expected token tab %d got %d", et.Tab, token.Tab)
		}
	}
}
