package Ex03TimeComplexityFrogJmp;

import (
    "testing"
	"github.com/stretchr/testify/assert"
)

func Test_FrogJmp_0_0_0_is_0(t *testing.T) {
    assert.Equal(t, 0, Solution(0, 0, 0));
}

func Test_FrogJmp_10_15_5_is_1(t *testing.T) {
    assert.Equal(t, 1, Solution(10, 15, 5));
}

func Test_FrogJmp_10_16_5_is_2(t *testing.T) {
    assert.Equal(t, 2, Solution(10, 16, 5));
}

func Test_FrogJmp_10_20_5_is_2(t *testing.T) {
    assert.Equal(t, 2, Solution(10, 20, 5));
}

func Test_FrogJmp_10_85_30_is_3(t *testing.T) {
    assert.Equal(t, 3, Solution(10, 85, 30));
}
