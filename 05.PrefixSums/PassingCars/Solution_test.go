package Ex05PrefixSumsPassingCars

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_PassingCars(t *testing.T) {
    assert.Equal(t, 5, Solution([]int {0, 1, 0, 1, 1}));
}