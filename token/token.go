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
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LTPAREN = "("
	RTPAREN = ")"
	LTBRACE = "{"
	RTBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	MINI     = "MINI" // Mini is same as let for binding
)
