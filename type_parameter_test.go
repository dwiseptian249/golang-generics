package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Tio")
	assert.Equal(t, "Tio", result)

	var resultNumber int = Length[int](25)
	assert.Equal(t, 25, resultNumber)

}
