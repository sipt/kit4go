package math_test

import (
	"testing"

	"github.com/sipt/algorithm/math"
	"github.com/stretchr/testify/assert"
)

func TestGcd(t *testing.T) {
	assert.Equal(t, math.Gcd(100, 1), 1)
	assert.Equal(t, math.Gcd(100, 2), 2)
	assert.Equal(t, math.Gcd(100, 50), 50)
	assert.Equal(t, math.Gcd(100, 45), 5)
}
