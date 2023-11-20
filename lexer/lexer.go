package lexer

import (
	"github.com/aaazlkm/go-interpreter/token"
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

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		// TODO = ! == !=の判定はdefaultでまとめるとよさそう
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tk.Literal = string(ch) + string(l.ch)
			tk.Type = token.EQ
		} else {
			tk = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case '(':
		tk = newToken(token.LPAREN, l.ch)
	case ')':
		tk = newToken(token.RPAREN, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '-':
		tk = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tk.Literal = string(ch) + string(l.ch)
			tk.Type = token.NOT_EQ
		} else {
			tk = newToken(token.BANG, l.ch)
		}
	case '*':
		tk = newToken(token.ASTERISK, l.ch)
	case '/':
		tk = newToken(token.SLASH, l.ch)
	case '<':
		tk = newToken(token.LT, l.ch)
	case '>':
		tk = newToken(token.GT, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdentifier()
			tk.Type = token.LookupIdent(tk.Literal)
			return tk // すでにreadIdentifierが呼ばれているため、ここでreturnする
		} else if isDigit(l.ch) {
			tk.Literal = l.readNumber()
			tk.Type = token.INT
			return tk // すでにreadNumberが呼ばれているため、ここでreturnする
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tk
}

// 新しいトークンを作成する
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// 空白をスキップする
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 1文字読み込む
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// 次の文字を見る
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// 英字を読み込む
func (l *Lexer) readIdentifier() string {
	prev := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[prev:l.position]
}

// 英字かどうか
func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_'
}

// intを読み込む
func (l *Lexer) readNumber() string {
	prev := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[prev:l.position]
}

// intかどうか
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
