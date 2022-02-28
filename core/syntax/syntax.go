package syntax

import (
	"regexp"

	"github.com/KlyuchnikovV/edigode/types"
)

type Parser struct {
	isWhitespace *regexp.Regexp
}

const (
	// Undefined    = "undefined"
	// Delimiter    = "delimiter"
	Symbols = "symbols"
	// OpeningBrace = "opening-brace"
	// ClosingBrace = "closing-brace"
	// Quote        = "quote"
	NewLine = "new-line"
	// Keyword      = "keyword"
	// Type        = "built-in-type"
	// BuiltInFunc = "built-in-func"
	// Comment     = "comment"
	// String      = "string"
	Whitespace = "whitespace"
)

func New() Parser {
	return Parser{
		isWhitespace: regexp.MustCompile(`\s+`),
	}
}

func (p Parser) Tokenize(text string) ([]types.Line, error) {
	var tokens, err = p.tokenize(text)
	if err != nil {
		return nil, err
	}

	var result = []types.Line{{}}
	for i, token := range tokens {
		tokens[i].ClassNames = token.GetClasses()
		if token.Classes[0].Name == NewLine {
			result = append(result, types.Line{})
		} else {
			result[len(result)-1] = append(result[len(result)-1], tokens[i])
		}
	}

	return result, nil
}

func (p Parser) tokenize(text string) ([]types.Token, error) {
	var (
		token = types.Token{
			Classes: []types.Class{types.NewClass(Symbols)},
		}
		result = make([]types.Token, 0)
	)
	for _, char := range text {
		var class = p.defineClass(char)
		if (token.Classes[0] == class || token.Classes[0].Name == "") && class.Name != NewLine {
			if token.Classes[0].Name == "" {
				token.Classes[0] = class
			}
			token.Value += string(char)
			continue
		}
		if len(token.Value) > 0 {
			result = append(result, token)
		}
		var str string
		if class.Name == NewLine {
			result = append(result, types.Token{
				Value:   "\n",
				Classes: types.Classes{class},
			})
			str = ""
			class.Name = ""
		} else {
			str = string(char)
		}

		token = types.Token{
			Value:   str,
			Classes: types.Classes{class},
		}
	}

	result = append(result, token)

	return result, nil
}

func (p Parser) defineClass(char rune) types.Class {
	if char == '\n' {
		return types.NewClass(NewLine)
	}

	if p.isWhitespace.Match([]byte{byte(char)}) {
		return types.NewClass(Whitespace)
	}

	return types.NewClass(Symbols)
}
