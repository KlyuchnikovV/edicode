package main

import (
	ctx "context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KlyuchnikovV/buffer"
)

func TestMain2(t *testing.T) {
	buf := *buffer.New(ctx.Background(), "")

	buf.Insert(1, 0, 'e')

	buf.Tree.String()
}

func TestNewFromString(t *testing.T) {
	var str = "package main\n\nimport \"fmt\"\n\nHello, world!"
	buf := *buffer.New(ctx.Background(), "", []byte(str)...)

	assert.Equal(t, str, buf.Tree.String())

	fmt.Printf("result: %s", buf.Tree.String())
}
