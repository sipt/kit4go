package util_test

import (
	"testing"

	"github.com/sipt/algorithm/util"
	"github.com/stretchr/testify/assert"
)

type TestInterface interface {
	Data()
}
type TestStruct struct{}

func (*TestStruct) Data() {}

func TestIsNil(t *testing.T) {
	var i TestInterface
	var s *TestStruct
	i = s
	assert.Equal(t, util.IsNil(i), true)
}
