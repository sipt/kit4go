package math_test

import (
	"testing"

	"github.com/sipt/algorithm/math"
	"github.com/stretchr/testify/assert"
)

func TestGcd(t *testing.T) {
	assert.Equal(t, math.Gcd(1000000, 1), 1)
	assert.Equal(t, math.Gcd(1000000, 2), 2)
	assert.Equal(t, math.Gcd(1000000, 50), 50)
	assert.Equal(t, math.Gcd(1000000, 45), 5)
}
