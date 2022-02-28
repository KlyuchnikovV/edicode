package types

import (
	"strings"

	"github.com/KlyuchnikovV/stack"
)

type Class struct {
	Name string `json:"name"`

	Trigger    string `json:"trigger"`
	Stopper    string `json:"stopper"`
	Multiline  bool   `json:"multiline"`
	Multitoken bool   `json:"multitoken"`
	Grouping   string `json:"grouping"`
}

func NewClass(name string) Class {
	return Class{
		Name: name,
	}
}

type Classes []Class

func (c Classes) String() string {
	var result = make([]string, len(c))
	for i, class := range c {
		result[i] = class.Name
	}
	return strings.Join(result, "|")
}

func (c Classes) Contains(class Class) bool {
	for _, cl := range c {
		if cl == class {
			return true
		}
	}
	return false
}

func (c *Classes) Append(classes ClassesStack) {
	*c = append(*c, classes.ToSlice()...)
}

type ClassesStack struct {
	stack stack.Stack
}

func NewClassesStack() *ClassesStack {
	return &ClassesStack{
		stack: *stack.New(5),
	}
}

func (st *ClassesStack) Push(classes ...Class) {
	for _, class := range classes {
		st.stack.Push(class)
	}
}

func (st *ClassesStack) Peek() (*Class, bool) {
	v, ok := st.stack.Peek()
	if !ok {
		return nil, false
	}
	class, ok := v.(Class)
	return &class, ok
}

func (st *ClassesStack) Pop() {
	st.stack.Pop()
}

func (st *ClassesStack) PopN(n int) {
	st.stack.PopN(n)
}

func (st *ClassesStack) Size() int {
	return st.stack.Size()
}

func (st *ClassesStack) ToSlice() []Class {
	var (
		slice  = st.stack.ToSlice()
		result = make([]Class, len(slice))
	)
	for i := range slice {
		result[i] = slice[i].(Class)
	}
	return result
}

type BuiltIn struct {
	Keyword []Class
}
