package types

import "fmt"

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
