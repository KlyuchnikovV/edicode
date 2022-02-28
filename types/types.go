package types

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
		Value      string   `json:"value"`
		Classes    Classes  `json:"-"`
		ClassNames []string `json:"classes"`
	}

	Text []Line
	Line []Token
)

func (t Token) GetClasses() []string {
	var classes = make([]string, len(t.Classes))
	for i, class := range t.Classes {
		classes[i] = class.Name
	}
	return classes
}
