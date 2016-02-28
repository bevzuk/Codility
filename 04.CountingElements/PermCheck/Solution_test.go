package Ex04CountingElementsPermCheck;

import (
    "testing"
	"github.com/stretchr/testify/assert"
)

func Test_PermCheck_Passed_For_1_2_3(t *testing.T) {
    assert.Equal(t, 1, Solution([]int {1, 2, 3}));
}

func Test_PermCheck_Failed_For_1_2_4(t *testing.T) {
    assert.Equal(t, 0, Solution([]int {1, 2, 4}));
}
