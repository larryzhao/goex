package intsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {}

func TestInclude(t *testing.T) {
	r := &Range{
		Start: 3,
		End:   10,
	}

	assert.Equal(t, true, r.Include(3))
	assert.Equal(t, true, r.Include(10))
	assert.Equal(t, false, r.Include(11))
}

func TestIncludeRange(t *testing.T) {
	r1 := &Range{
		Start: 3,
		End: 10,
	}

	r2 := &Range{
		Start: 12,
		End: 20,
	}

	r3 := &Range{
		Start: 4,
		End: 9,
	}

	r4 := &Range{
		Start: 7,
		End: 20,
	}

	assert.Equal(t, false, r1.IncludeRange(r2))
	assert.Equal(t, true, r1.IncludeRange(r3))
	assert.Equal(t, false, r1.IncludeRange(r4))
}

func TestIter(t *testing.T) {
	r1 := &Range{
		Start: 3,
		End: 10,
	}

	r1.Iter(func(idx int, val int) {
		assert.Equal(t, idx + 3, val)
	})
}
