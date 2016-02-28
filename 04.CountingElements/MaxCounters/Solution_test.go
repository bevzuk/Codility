package Ex04CoutingElementsMaxCounters

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_MaxCounters(t *testing.T) {
    assert.Equal(t, []int {1}, Solution(1, []int {1}));
}

func Test_MaxCounters_2(t *testing.T) {
    assert.Equal(t, []int {1, 2}, Solution(2, []int {1, 3, 2}));
}

func Test_Counters_CanIncrement(t *testing.T) {
    counters := NewCounters(2);
    counters.Apply(1);
    assert.Equal(t, []int {1, 0}, counters.GetCounters());
}