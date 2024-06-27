package json

import "bufio"

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	reader *bufio.Reader
}

const (
	TokenEOF TokenType = iota
	TokenLBrace
	TokenRBrace
	TokenColon
	TokenComma
	TokenString
	TokenNumber
	TokenBool
	TokenNull
)

type Parser struct {
	Lexer *Lexer
	Token Token
	Target interface{}
}
