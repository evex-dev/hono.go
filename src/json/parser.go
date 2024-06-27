package json

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func NewParser(lexer *Lexer, target interface{}) *Parser {
	return &Parser{Lexer: lexer, Target: target}
}

func (p *Parser) Parse() error {
	var err error
	p.Token, err = p.Lexer.NextToken()
	if err != nil {
		return err
	}
	return p.parseValue(reflect.ValueOf(p.Target).Elem())
}

func (p *Parser) parseValue(v reflect.Value) error {
	switch p.Token.Type {
	case TokenLBrace:
		return p.parseObject(v)
	case TokenString:
		v.SetString(p.Token.Value)
		p.Token, _ = p.Lexer.NextToken()
		return nil
	case TokenNumber:
		return p.parseNumber(v)
	case TokenBool:
		v.SetBool(p.Token.Value == "true")
		p.Token, _ = p.Lexer.NextToken()
		return nil
	case TokenNull:
		v.Set(reflect.Zero(v.Type()))
		p.Token, _ = p.Lexer.NextToken()
		return nil
	default:
		return fmt.Errorf("unexpected token: %v", p.Token)
	}
}

func (p *Parser) parseObject(v reflect.Value) error {
	p.Token, _ = p.Lexer.NextToken()
	for p.Token.Type != TokenRBrace {
		if p.Token.Type != TokenString {
			return fmt.Errorf("expected string key, got: %v", p.Token)
		}
		key := p.Token.Value
		p.Token, _ = p.Lexer.NextToken()

		if p.Token.Type != TokenColon {
			return fmt.Errorf("expected colon after key, got: %v", p.Token)
		}
		p.Token, _ = p.Lexer.NextToken()

		field := p.findFieldByJSONtag(v, key)

		if !field.IsValid() {
			return fmt.Errorf("unknown field: %s", key)
		}
		if err := p.parseValue(field); err != nil {
			return err
		}

		if p.Token.Type == TokenColon {
			p.Token, _ = p.Lexer.NextToken()
		}
	}
	p.Token, _ = p.Lexer.NextToken()

	return nil
}

func (p *Parser) parseNumber(v reflect.Value) error {
	if v.Kind() == reflect.Int {
		n, err := strconv.Atoi(p.Token.Value)
		if err != nil {
			return err
		}
		v.SetInt(int64(n))
	} else if v.Kind() == reflect.Float64 {
		n, err := strconv.ParseFloat(p.Token.Value, 64)
		if err != nil {
			return err
		}
		v.SetFloat(n)
	}
	p.Token, _ = p.Lexer.NextToken()
	return nil
}

func (p *Parser) findFieldByJSONtag(v reflect.Value, tag string) reflect.Value {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == tag {
			return v.Field(i)
		}
	}
	return reflect.Value{}
}

func ParseJSON(value string, target interface{}) error {
	reader := strings.NewReader(value)
	lexer := NewLexer(reader)
	parser := NewParser(lexer, target)
	return parser.Parse()
}
