/* LAAIT (pronounce same as lite) Programming Language
   It is C like syntax interpreted, language contains

   --- Lower Level Features --
   Integer and Booleans
   String
   Array
   Arithmetic Expressions
   Built-in Functions

   -- Higher Level Features --
   higher level feature as
   Variable Binding
   First-Class and Higher Order Functions
   Closures


   Vim is used as Text Editor and Pop OS 21.04 is the host os.
   Starting Date is Thu Jul 22 07:20:16 PM IST 2021

*/

/*  This is token package, use to classification of tokens used */

package token

type TokenType string /* TokenType is a string type  which helps us in debugging and
   in making difference between them*/

type Token struct {
	Type    TokenType
	Literal string
}

/* Token types in LAAIT */
const (
	FILEEND = "EOF"
	ILLEGAL = "ILLEGAL" // ILLEGAL signifies token character we donot know about

	// Identifiers and Literals
	NOUN = "NOUN" // Noun is alias for Identifiers
	INT  = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
    QUOTE    = "\""

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LTPAREN = "("
	RTPAREN = ")"
	LTBRACE = "{"
	RTBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	BIND     = "BIND" // bind is same as let for binding
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	NOT      = "NOT"
	OR       = "OR"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"bind":   BIND,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupNoun(noun string) TokenType {
	if tok, ok := keywords[noun]; ok {
		return tok
	}
	return NOUN
}
