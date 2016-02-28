package Ex05PrefixSumsCountDiv

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_CountDiv(t *testing.T) {
    assert.Equal(t, 3, Solution(6, 11, 2));
    assert.Equal(t, 4, Solution(6, 12, 2));
    assert.Equal(t, 1, Solution(5, 7, 2));
    assert.Equal(t, 0, Solution(1, 1, 11));
    assert.Equal(t, 1, Solution(0, 0, 11));
    assert.Equal(t, 1, Solution(10, 10, 5));
    assert.Equal(t, 0, Solution(10, 10, 7));
    assert.Equal(t, 0, Solution(10, 10, 20));
}

// [7-11], 10 -> [0, 19] -> 1
// [7-20], 10 -> [0, 29] -> 2
// [10-20], 10 -> [0, 29] -> 2
// [11-20], 10 -> [10, 29] -> 1
// [11-19], 10 -> [10, 19] -> 0
// [6-11], 2 -> [4, 11] -> 3
// [6-12], 2 -> [4, 12] -> 4
// [5-7], 2 -> [4, 7] -> 1
// [1-1], 11 -> [0, 1] -> 0
// [0-0], 11 -> [-11, 0] -> 1