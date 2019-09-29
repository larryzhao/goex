package stringsx

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncate(t *testing.T) {
	a := "我我我爱爱爱你你你"
	assert.Equal(t, Truncate(a, 100, ""), a)
	assert.Equal(t, Truncate(a, 5, ""), "我我我爱爱")
	assert.Equal(t, Truncate(a, 5, "..."), "我我...")
	assert.Equal(t, Truncate(a, 6, "**"), "我我我爱**")
}

func BenchmarkTruncate(b *testing.B) {
	sb := &strings.Builder{}
	for i := 0; i < 150; i++ {
		fmt.Fprint(sb, "我")
	}

	testString := sb.String()

	for i := 0; i < b.N; i++ {
		Truncate(testString, 100, "...")
	}
}
