package lexer 

import "laait/token"

type Lexer struct{
    input string
    position int
    readPosition int 
    ch byte
}

func New(input string) *Lexer{
    l := &Lexer{input:input}
    l.readChar()
    return l
}

func (l *Lexer) readChar(){
    if l.readPosition >= len(l.input){
        l.ch = 0
    }else{
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition+=1
}

func (l *Lexer) NextToken() token.Token{
    var tok token.Token

    switch l.ch{
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LTPAREN, l.ch)
    case ')':
        tok = newToken(token.RTPAREN, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case '{':
        tok = newToken(token.LTBRACE, l.ch)
    case '}':
        tok = newToken(token.RTBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.FILEEND
    }
    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
    return token.Token{Type: tokenType, Literal : string(ch)}
}
