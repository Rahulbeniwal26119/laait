package lexer

import (
	"laait/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= utf8.RuneCountInString(l.input) {
		l.ch = 0
	} else {
		// writing to support UTF-8 characters
		l.ch = []rune(l.input)[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LTPAREN, l.ch)
	case ')':
		tok = newToken(token.RTPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LTBRACE, l.ch)
	case '}':
		tok = newToken(token.RTBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.FILEEND
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LTBRACKET, l.ch)
	case ']':
		tok = newToken(token.RTBRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	default:
		if unicode.IsLetter(l.ch) {
			tok.Literal = l.readNoun()
			tok.Type = token.LookupNoun(tok.Literal)
			return tok
		} else if unicode.IsDigit(l.ch) { // handle the interger numbers
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// read the string

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

// read the integer literal
func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNoun function reads the indentifier or nouns until, It encounter an non-letter-character

func (l *Lexer) readNoun() string {
	position := l.position
	for unicode.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Decide which character are allowed as identifier name, use either ! or ?,
// but i am planning ! for factorial function in uncoming code

// return the new token
func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// skip white spaces from the the source code while parsing

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
