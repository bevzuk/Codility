package Ex04CountingElementsMissingInteger

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
    assert.Equal(t, 2, Solution([]int {1, 3}));
}