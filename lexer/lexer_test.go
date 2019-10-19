package lexer

import (
	"testing"

	"github.com/kstola2/monkey_interpreter/mytoken"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
		};
		
		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;
		if (5 < 10) {
			return true;
		} else {
			return false;
		}
		
		10 == 10
		10 != 9
		`

	tests := []struct {
		expectedType    mytoken.TokenType
		expectedLiteral string
	}{
		{mytoken.LET, "let"},
		{mytoken.IDENT, "five"},
		{mytoken.ASSIGN, "="},
		{mytoken.INT, "5"},
		{mytoken.SEMICOLON, ";"},
		{mytoken.LET, "let"},
		{mytoken.IDENT, "ten"},
		{mytoken.ASSIGN, "="},
		{mytoken.INT, "10"},
		{mytoken.SEMICOLON, ";"},
		{mytoken.LET, "let"},
		{mytoken.IDENT, "add"},
		{mytoken.ASSIGN, "="},
		{mytoken.FUNCTION, "fn"},
		{mytoken.LPAREN, "("},
		{mytoken.IDENT, "x"},
		{mytoken.COMMA, ","},
		{mytoken.IDENT, "y"},
		{mytoken.RPAREN, ")"},
		{mytoken.LBRACE, "{"},
		{mytoken.IDENT, "x"},
		{mytoken.PLUS, "+"},
		{mytoken.IDENT, "y"},
		{mytoken.SEMICOLON, ";"},
		{mytoken.RBRACE, "}"},
		{mytoken.SEMICOLON, ";"},
		{mytoken.LET, "let"},
		{mytoken.IDENT, "result"},
		{mytoken.ASSIGN, "="},
		{mytoken.IDENT, "add"},
		{mytoken.LPAREN, "("},
		{mytoken.IDENT, "five"},
		{mytoken.COMMA, ","},
		{mytoken.IDENT, "ten"},
		{mytoken.RPAREN, ")"},
		{mytoken.SEMICOLON, ";"},
		{mytoken.BANG, "!"},
		{mytoken.MINUS, "-"},
		{mytoken.SLASH, "/"},
		{mytoken.ASTERISK, "*"},
		{mytoken.INT, "5"},
		{mytoken.SEMICOLON, ";"},
		{mytoken.INT, "5"},
		{mytoken.LT, "<"},
		{mytoken.INT, "10"},
		{mytoken.GT, ">"},
		{mytoken.INT, "5"},
		{mytoken.SEMICOLON, ";"},
		{mytoken.IF, "if"},
		{mytoken.LPAREN, "("},
		{mytoken.INT, "5"},
		{mytoken.LT, "<"},
		{mytoken.INT, "10"},
		{mytoken.RPAREN, ")"},
		{mytoken.LBRACE, "{"},
		{mytoken.RETURN, "return"},
		{mytoken.TRUE, "true"},
		{mytoken.SEMICOLON, ";"},
		{mytoken.RBRACE, "}"},
		{mytoken.ELSE, "else"},
		{mytoken.LBRACE, "{"},
		{mytoken.RETURN, "return"},
		{mytoken.FALSE, "false"},
		{mytoken.SEMICOLON, ";"},
		{mytoken.RBRACE, "}"},
		{mytoken.INT, "10"},
		{mytoken.EQ, "=="},
		{mytoken.INT, "10"},
		{mytoken.INT, "10"},
		{mytoken.NOT_EQ, "!="},
		{mytoken.INT, "9"},
		{mytoken.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
