package actions

type Action interface {
	Filter(string) bool
	Name() string
	Do(string) error
}

