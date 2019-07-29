package lexer

import (
	"github.com/kstola2/monkey_interpreter/mytoken"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() mytoken.Token {
	var tok mytoken.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peerChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = mytoken.Token{Type: mytoken.EQ, Literal: literal}
		} else {
			tok = newToken(mytoken.ASSIGN, l.ch)
		}
	case '-':
		tok = newToken(mytoken.MINUS, l.ch)
	case '!':
		if l.peerChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = mytoken.Token{Type: mytoken.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(mytoken.BANG, l.ch)
		}
	case '/':
		tok = newToken(mytoken.SLASH, l.ch)
	case '*':
		tok = newToken(mytoken.ASTERISK, l.ch)
	case '<':
		tok = newToken(mytoken.LT, l.ch)
	case '>':
		tok = newToken(mytoken.GT, l.ch)
	case ';':
		tok = newToken(mytoken.SEMICOLON, l.ch)
	case '(':
		tok = newToken(mytoken.LPAREN, l.ch)
	case ')':
		tok = newToken(mytoken.RPAREN, l.ch)
	case ',':
		tok = newToken(mytoken.COMMA, l.ch)
	case '+':
		tok = newToken(mytoken.PLUS, l.ch)
	case '{':
		tok = newToken(mytoken.LBRACE, l.ch)
	case '}':
		tok = newToken(mytoken.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = mytoken.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = mytoken.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = mytoken.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(mytoken.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType mytoken.TokenType, ch byte) mytoken.Token {
	return mytoken.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peerChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
