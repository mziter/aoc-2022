package pointset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointSet(t *testing.T) {
	ps := New(256)
	ps.Insert(3, 4)
	ps.Insert(4, 3)
	ps.Insert(14, 223)

	assert.True(t, ps.Contains(3, 4))
	assert.True(t, ps.Contains(4, 3))
	assert.True(t, ps.Contains(14, 223))
}
