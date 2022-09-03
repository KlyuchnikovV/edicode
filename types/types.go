package types

import (
	"github.com/KlyuchnikovV/simple_buffer/selection"
)

type Core interface {
	GetBuffer(string) (*BufferData, error)

	OnInit(string, func(...interface{}))
	Emit(string, string, interface{})
	EmitTimed(string, string, int64, interface{})
}

type BufferData struct {
	Buffer    string `json:"buffer"`
	Symbol    int    `json:"symbol"`
	Lines     []Line `json:"lines"`
	Timestamp int64  `json:"timestamp"`
}

type (
	Token struct {
		Offset     int      `json:"offset"`
		Value      string   `json:"value"`
		ClassNames []string `json:"classes"`
		Classes    Classes  `json:"-"`
	}

	Text []Line
	Line []Token
)

func (t Token) NextOffset() int {
	if t.Classes.Contains(NewClass("new-line")) {
		return 0
	}
	return t.Offset + len(t.Value)
}

func (t Token) GetClasses() []string {
	var classes = make([]string, len(t.Classes))
	for i, class := range t.Classes {
		classes[i] = class.Name
	}
	return classes
}

type GetLineLengthRequest struct {
	Buffer string `json:"buffer"`
	Line   int    `json:"line"`
}

type GetCaretResponse struct {
	Buffer string          `json:"buffer"`
	Start  selection.Caret `json:"start"`
	End    selection.Caret `json:"end"`
}

type CaretMovedEvent struct {
	Buffer string          `json:"buffer"`
	Start  selection.Caret `json:"start"`
	End    selection.Caret `json:"end"`
}

type MouseEvent struct {
	Buffer string `json:"buffer"`
	Line   int    `json:"line"`
	Offset int    `json:"offset"`
}

type ActionParams struct {
	Action string `json:"action"`
	Param  string `json:"param"`
}
