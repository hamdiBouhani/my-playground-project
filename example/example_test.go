package example

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {

	// Variadic Functions
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// Closures
	nextInt := intSeq()
	assert.Equal(t, 1, nextInt())
	assert.Equal(t, 2, nextInt())
	assert.Equal(t, 3, nextInt())

	newInts := intSeq()
	assert.Equal(t, 1, newInts())

	//Recursion
	assert.Equal(t, 5040, fact(7))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	assert.Equal(t, 13, fib(7))
}

func TestStructEmbedding(t *testing.T) {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	type describer interface {
		describe() string
	}

	var d describer = co
	str := fmt.Sprintf("describer: %s", d.describe())
	assert.Equal(t, "describer: base with num=1", str)
}

func TestGenerics(t *testing.T) {
	// var m = map[int]string{1: "2", 2: "4", 4: "8"}
	// r := MapKeys(m)
	// j := 0
	// for k := range m {
	// 	assert.Equal(t, k, r[j])
	// 	j++
	// }

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
