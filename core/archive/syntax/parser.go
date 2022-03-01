package syntax

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type SyntaxParser struct {
	Symbols map[string][]string `json:"symbols"`
	Types   []Class             `json:"types"`
	// BuiltIn BuiltIn             `json:"built-in"`
}

func New(path string) SyntaxParser {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var parser SyntaxParser
	if err := json.Unmarshal(bytes, &parser); err != nil {
		panic(err)
	}

	return parser
}

func (p SyntaxParser) ParseText(text string) ([]Line, error) {
	tokens, err := p.Tokenize(text)
	if err != nil {
		return nil, err
	}
	tokens = p.ProcessTokens(tokens, *NewClassesStack())

	var result = []Line{{}}
	for i, token := range tokens {
		if token.Classes[0].Name == newLine {
			result = append(result, NewLine(nil))
		} else {
			result[len(result)-1] = append(result[len(result)-1], tokens[i])
		}
	}

	return result, nil
}

func (p *SyntaxParser) Tokenize(text string) ([]Token, error) {
	var (
		token = Token{
			Classes: []Class{NewClass(Symbols)},
		}
		result = make([]Token, 0)
	)
	for _, char := range text {
		var class = p.defineClass(char)
		if token.Classes[0] == class && class.Name != newLine {
			token.Value += string(char)
			continue
		}
		if len(token.Value) > 0 {
			result = append(result, token)
		}
		token = NewToken(string(char), class)
	}

	if len(token.Value) != 0 {
		result = append(result, token)
	}
	return result, nil
}

func (p *SyntaxParser) ProcessTokens(tokens []Token, context ClassesStack) []Token {
	var (
		immutableContextPart = context.Size()
		result               []Token
	)

	for i := 0; i < len(tokens); i++ {
		tokens[i].Classes = []Class{tokens[i].Classes[0]}

		var (
			classes = p.DefineToken(tokens[i])
			class   = classes[len(classes)-1]
		)

		if immutableContextPart == context.Size() {
			cl, ok := context.Peek()
			if ok && *cl == class && tokens[i].Value == class.Stopper {
				tokens[i].Classes.Append(context)
				result = append(result, tokens[i])
				break
			}
		}

		context.Push(classes...)

		// append classes from stack
		tokens[i].Classes.Append(context)
		result = append(result, tokens[i])

		if len(classes) > 0 {
			context.PopN(len(classes) - 1)
		}

		if class.Multiline || class.Multitoken {
			result = append(result, p.ProcessTokens(tokens[i+1:], context)...)
			i = len(result) - 1
		}

		context.Pop()
		if class.Name == newLine {
			cl, ok := context.Peek()
			if ok && !cl.Multiline {
				break
			}
		}

		// if immutableContextPart == context.Size() {
		// 	cl, ok := context.Peek()
		// 	if ok && *cl == class && tokens[i].Value == class.Stopper {
		// 		break
		// 	}
		// }

		// if class

		// switch class.Name {
		// case newLine:
		// 	if context.Size() > immutableContextPart {
		// 		context.Pop()
		// 	}
		// }

		// switch tokens[i].Classes[0].Name {
		// case Delimiter:
		// 	switch tokens[i].Value {
		// 	case "//":
		// 		context.Push(NewClass(String))
		// 		t := p.ProcessTokens(tokens[i:p.FindElementPosition(tokens[i:], "\n")], context)
		// 		for j := range t {
		// 			tokens[i] = t[j]
		// 		}
		// 		i = len(result) - 1
		// 		context.Pop()
		// 		continue
		// 	}
		// case Symbols:
		// 	switch tokens[i].Value {
		// 	case "package", "func", "var", "import", "for", "switch", "type", "return", "if", "else", "case":
		// 		tokens[i].Classes = append(tokens[i].Classes, NewClass(Keyword))
		// 		context.Push(NewClass(tokens[i].Value))
		// 	case "int", "bool", "string", "struct", "interface", "error", "nil":
		// 		tokens[i].Classes = append(tokens[i].Classes, NewClass(Type))
		// 	case "append", "len", "cap", "panic", "go", "defer":
		// 		tokens[i].Classes = append(tokens[i].Classes, NewClass(BuiltInFunc))
		// 		context.Push(NewClass(tokens[i].Value))
		// 	}
		// case newLine:
		// 	if context.Size() > immutableContextPart {
		// 		context.Pop()
		// 	}
		// case Quote:
		// context.Push(NewClass(String))
		// t, err := p.ProcessTokens(p.PairQuote(tokens[i:]), context)
		// if err != nil {
		// 	return nil, err
		// }
		// for j := range t {
		// 	tokens[i] = t[j]
		// }
		// i = len(result) - 1
		// context.Pop()
		// continue
		// case OpeningBrace:
		// 	t, err := p.ProcessTokens(p.PairBrace(tokens[i:]), context)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	for j := range t {
		// 		tokens[i] = t[j]
		// 	}
		// 	i = len(result) - 1
		// 	continue
		// }
	}

	return result
}

func (p *SyntaxParser) FindElementPosition(tokens []Token, value string) int {
	for i, token := range tokens {
		if string(token.Value) == value {
			return i
		}
	}
	return -1
}

func (p *SyntaxParser) DefineToken(token Token) []Class {
	if token.Classes[0] == NewClass(newLine) {
		return []Class{token.Classes[0]}
	}
	for _, class := range p.Types {
		if class.Trigger == token.Value {
			return []Class{class}
		}
	}

	// for _, class := range p.BuiltIn.Keyword {
	// 	if class.Name == token.Value {
	// 		return []Class{NewClass("keyword"), class}
	// 	}
	// }
	return []Class{NewClass(Undefined)}
}

func (p *SyntaxParser) defineClass(char rune) Class {
	if char == '\n' {
		return NewClass(newLine)
	}

	for class, tests := range p.Symbols {
		for _, test := range tests {
			if test == string(char) {
				return NewClass(class)
			}
		}
	}

	return NewClass(Symbols)
}
