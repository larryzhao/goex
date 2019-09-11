package stringsex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncate(t *testing.T) {
	a := "helloworld"
	assert.Equal(t, Truncate(a, 100), a)
	assert.Equal(t, Truncate(a, 5), "hello")
}
