package syntax

type Trigger struct {
	Trigger string `json:"trigger"`
}

type Text []Line

type Line []Token

func NewLine(tokens []Token) Line {
	return tokens
}

const (
	Undefined    = "undefined"
	Delimiter    = "delimiter"
	Symbols      = "symbols"
	OpeningBrace = "opening-brace"
	ClosingBrace = "closing-brace"
	Quote        = "quote"
	newLine      = "new-line"
	// Keyword      = "keyword"
	Type        = "built-in-type"
	BuiltInFunc = "built-in-func"
	Comment     = "comment"
	String      = "string"
	Whitespace  = "whitespace"
)

func bracesHasSameType(left, right Token) bool {
	if left.Classes[0].Name != OpeningBrace {
		return false
	}
	if right.Classes[0].Name != ClosingBrace {
		return false
	}
	switch {
	case left.Value[0] == '(' && right.Value[0] == ')',
		left.Value[0] == '[' && right.Value[0] == ']',
		left.Value[0] == '{' && right.Value[0] == '}':
		return true
	}
	return false
}
