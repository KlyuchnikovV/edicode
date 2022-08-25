package types

import "fmt"

type KeyboardEvent struct {
	Buffer string `json:"buffer"`
	Key    string `json:"key"`
	Alt    bool   `json:"alt"`
	Ctrl   bool   `json:"ctrl"`
	Shift  bool   `json:"shift"`
	Meta   bool   `json:"meta"`

	StartLine   int `json:"startLine"`
	StartOffset int `json:"startOffset"`
	EndLine     int `json:"endLine"`
	EndOffset   int `json:"endOffset"`

	// Line   int `json:"line"`
	Offset int `json:"offset"`
}

func (e KeyboardEvent) String() string {
	return fmt.Sprintf("for buffer '%s' key: %s, alt: %t, ctrl: %t, shift: %t",
		e.Buffer, e.Key, e.Alt, e.Ctrl, e.Shift,
	)
}

type SelectEvent struct {
	Buffer string `json:"buffer"`
	Line   *int   `json:"line,omitempty"`
	Symbol *int   `json:"symbol,omitempty"`
}

func (e SelectEvent) String() string {
	var result = fmt.Sprintf("for buffer '%s'", e.Buffer)

	if e.Line != nil {
		result = fmt.Sprintf("%s offset %d", result, *e.Line)
	}
	if e.Symbol != nil {
		result = fmt.Sprintf("%s offset %d", result, *e.Symbol)
	}

	return result
}

// type SelectLineEvent struct {
// 	Buffer string `json:"buffer"`
// 	Line   int    `json:"line"`
// }
