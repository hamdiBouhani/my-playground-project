package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastInt(t *testing.T) {

	data := []int{1, 2, 3}
	assert.Equal(t, 3, Last[int](data))

	data2 := []string{"a", "b", "c"}
	assert.Equal(t, "c", Last[string](data2))
}
