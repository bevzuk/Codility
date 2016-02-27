package Ex01IterationsBinaryGap;

import (
    "strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_BinaryGapOf_0_Is_0(t *testing.T) {
	assert.Equal(t, 0, Solution(0));
}

func Test_BinaryGapOf_1_Is_0(t *testing.T) {
	assert.Equal(t, 0, Solution(1));
}

func Test_BinaryGapOf_101_Is_1(t *testing.T) {
	assert.Equal(t, 1, SolutionStr("101"));
}

func Test_BinaryGapOf_0001001000_Is_2(t *testing.T) {
	assert.Equal(t, 2, SolutionStr("0001001000"));
}

func Test_BinaryGapOf_00010010001_Is_3(t *testing.T) {
	assert.Equal(t, 3, SolutionStr("00010010001"));
}

func Test_BinaryGapOf_100100000101_Is_5(t *testing.T) {
	assert.Equal(t, 5, SolutionStr("100100000101"));
}

func SolutionStr(s string) int {
    if i, err := strconv.ParseInt(s, 2, 64); err == nil {
        return Solution(int(i));
    }
    return 0;
}