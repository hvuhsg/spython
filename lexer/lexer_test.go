package lexer

import (
	"testing"

	"github.com/hvuhsg/spython/token"
)

func TestLexer(t *testing.T) {
	lexer := New("+ = \n \t (==)")

	expectedTokens := []token.Token{
		token.New(token.PLUS, "+", 0, 0, 0),
		token.New(token.ASSIGN, "=", 0, 2, 0),
		token.New(token.LEFT_PARAN, "(", 1, 3, 1),
		token.New(token.EQUALS, "==", 1, 4, 1),
		token.New(token.RIGHT_PARAN, ")", 1, 6, 1),
		token.New(token.EOF, "EOF", 1, 7, 1),
	}

	for _, et := range expectedTokens {
		token := lexer.NextToken()

		if token.Type != et.Type {
			t.Fatalf("Expected token type %d got %d", et.Type, token.Type)
		}

		if token.Value != et.Value {
			t.Fatalf("Expected token value '%s' got '%s'", et.Value, token.Value)
		}

		if token.Row != et.Row {
			t.Fatalf("Expected token row %d got %d", et.Row, token.Row)
		}

		if token.Col != et.Col {
			t.Fatalf("Expected token col %d got %d", et.Col, token.Col)
		}

		if token.Row != et.Row {
			t.Fatalf("Expected token tab %d got %d", et.Tab, token.Tab)
		}
	}
}
