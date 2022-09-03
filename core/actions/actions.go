package actions

import (
	"strings"
)

// Find by part of string
// Find by string
//

var Actions = []Action{
	&SaveFile{},
	&OpenFile{},
}

type SaveFile struct {
	do func(string) error
}

func NewSaveFile(do func(string) error) *SaveFile {
	return &SaveFile{
		do: do,
	}
}

func (action SaveFile) Filter(f string) bool {
	var value = "save"

	for i, r := range strings.ToLower(f) {
		if len(value) <= i {
			return false
		}

		if rune(value[i]) != r {
			return false
		}
	}

	return true
}

func (action SaveFile) Name() string {
	return "Save"
}

func (action *SaveFile) Do(s string) error {
	return action.do(s)
}

type OpenFile struct {
}

func (action OpenFile) Filter(f string) bool {
	var value = "open"

	for i, r := range strings.ToLower(f) {
		if len(value) <= i {
			return false
		}

		if rune(value[i]) != r {
			return false
		}
	}

	return true
}

func (action OpenFile) Name() string {
	return "Open"
}

func (action *OpenFile) Do(string) error {
	return nil
}
