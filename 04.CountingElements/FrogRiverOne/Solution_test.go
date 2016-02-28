package Ex04CountingElementsFrogRiverOne;

import (
    "testing"
	"github.com/stretchr/testify/assert"
)

func Test_FrogRiverOne_CanCrossRiver_DirectOrder(t *testing.T) {
    assert.Equal(t, 2, Solution(3, []int {1, 2, 3}));
}

func Test_FrogRiverOne_CanCrossRiver_Mixed(t *testing.T) {
    assert.Equal(t, 6, Solution(5, []int {1, 3, 1, 4, 2, 3, 5, 4}));
}

func Test_FrogRiverOne_CanNotCrossRiver_WithGap(t *testing.T) {
    assert.Equal(t, -1, Solution(3, []int {1, 1, 2}));
}
