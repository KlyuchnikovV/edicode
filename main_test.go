package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	buffer "github.com/KlyuchnikovV/simple_buffer"
)

func TestMain2(t *testing.T) {
	buf := *buffer.New()

	buf.Insert(1, 0, 'e')

	fmt.Printf(buf.String())
}

func TestNewFromString(t *testing.T) {
	var str = "package main\n\nimport \"fmt\"\n\nHello, world!"
	buf := *buffer.New([]rune(str)...)

	assert.Equal(t, str, buf.String())

	fmt.Printf("result: %s", buf.String())
}
