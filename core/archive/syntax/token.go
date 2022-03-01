package syntax

type Token struct {
	Value   string  `json:"value"`
	Classes Classes `json:"classes"`
}

func NewToken(runes string, class ...Class) Token {
	return Token{
		Value:   runes,
		Classes: class,
	}
}

func (token Token) Length() int {
	return len(token.Value)
}

func (token *Token) Insert(position int, runes string) {
	var temp = token.Value

	token.Value = temp[:position] + runes + temp[position:]
}

func (token *Token) Remove(position int) {
	if position < 0 || position >= len(token.Value) {
		return
	}
	var temp = token.Value

	token.Value = temp[:position] + temp[position+1:]
}

func (token Token) Split(position int) (Token, Token) {
	var temp = token.Value
	return NewToken(temp[:position], token.Classes...), NewToken(temp[position:], token.Classes...)
}
