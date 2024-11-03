package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	Identifiers = "Identifiers"
	Int         = "Int"
	String      = "String"
	Type        = "Type"
	// Operators
	Assign = "="
	Plus   = "+"
	Minus  = "-"

	//Delimiters
	Comma     = ","
	Semicolon = ";"
	LParen    = "("
	RParen    = ")"
	LBrace    = "{"
	RBrace    = "}"
	Dot       = "."
	// Keywords
	Function = "MrviFn"
	Var      = "MrviVar"
	Return   = "MrviReturn"
)
