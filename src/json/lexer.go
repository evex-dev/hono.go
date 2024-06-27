package json

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{reader: bufio.NewReader(reader)}
}

func (l *Lexer) NextToken() (Token, error) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return Token{Type: TokenEOF}, nil
			}
			return Token{}, err
		}
		switch {
		case unicode.IsSpace(r):
			continue
		case r == '{':
			return Token{Type: TokenLBrace}, nil
		case r == '}':
			return Token{Type: TokenRBrace}, nil
		case r == ':':
			return Token{Type: TokenColon}, nil
		case r == ',':
			return Token{Type: TokenComma}, nil
		case r == '"':
			return l.readString()
		case unicode.IsDigit(r) || r == '-':
			return l.readNumber(r)
		case r == 't' || r == 'f':
			return l.readBool(r)
		case r == 'n':
			return l.readNull()
		default:
			return Token{}, fmt.Errorf("unexpected character: %v", r)
		}
	}
}

func (l *Lexer) readString() (Token, error) {
	var sb strings.Builder
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			return Token{}, err
		}
		if r == '"' {
			break
		}
		sb.WriteRune(r)
	}
	return Token{Type: TokenString, Value: sb.String()}, nil
}

func (l *Lexer) readNumber(initial rune) (Token, error) {
	var sb strings.Builder
	sb.WriteRune(initial)
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return Token{Type: TokenEOF}, nil
			}
			return Token{}, err
		}
		if !unicode.IsDigit(r) && r != '.' && r != 'e' && r != 'E' && r != '+' && r != '-' {
			l.reader.UnreadRune()
			break
		}

		sb.WriteRune(r)
	}
	return Token{Type: TokenNumber, Value: sb.String()}, nil
}

func (l *Lexer) readBool(initial rune) (Token, error) {
	var expected string
	if initial == 't' {
		expected = "rue"
	} else {
		expected = "alse"
	}

	for _, r := range expected {
		r2, _, err := l.reader.ReadRune()
		if err != nil {
			return Token{}, err
		}
		if r != r2 {
			return Token{}, fmt.Errorf("unexpected character in boolean")
		}
	}

	return Token{Type: TokenBool, Value: string(initial) + expected}, nil
}

func (l *Lexer) readNull() (Token, error) {
	expected := "ull"
	for _, r := range expected {
		r2, _, err := l.reader.ReadRune()
		if err != nil {
			return Token{}, err
		}

		if r != r2 {
			return Token{}, fmt.Errorf("unexpected character in null")
		}
	}

	return Token{Type: TokenNull, Value: "null"}, nil
}
