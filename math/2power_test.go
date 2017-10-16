package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs2Power(t *testing.T) {
	var value = 100
	assert.Equal(t, Is2Power(value), false, "100是2的乘方？正确结果：false，你的结果: true")
	value = 1024
	assert.Equal(t, Is2Power(value), true, "1024是2的乘方？正确结果：false，你的结果: true")
}
