package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOG     = "EOF"
	// Identifiers + literals

	IDENT = "IDENT" // add, foobar, x, y,
	INT   = "INT"

	//OPS
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
