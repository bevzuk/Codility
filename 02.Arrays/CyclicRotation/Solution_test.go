package Ex02ArraysCyclicRotation

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_CyclicRotation_0_DoesNothing(t *testing.T) {
    array := []int{1, 2, 3};
    cycledArray := Solution(array, 0)
	assert.Equal(t, []int{1, 2, 3}, cycledArray);
}

func Test_CyclicRotation_1_CyclesRightTo1(t *testing.T) {
    array := []int{1, 2, 3};
    cycledArray := Solution(array, 1)
	assert.Equal(t, []int{3, 1, 2}, cycledArray);
}

func Test_CyclicRotation_2_CyclesRightTo2(t *testing.T) {
    array := []int{1, 2, 3, 4, 5};
    cycledArray := Solution(array, 2)
	assert.Equal(t, []int{4, 5, 1, 2, 3}, cycledArray);
}