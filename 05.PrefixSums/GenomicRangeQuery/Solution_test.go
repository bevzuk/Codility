package Ex03PrefixSumsGenomicRangeQuery;

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_GenomicRangeQuery_Sample1(t *testing.T) {
    result := Solution("CAGCCTA", []int{2, 5, 0}, []int {4, 5, 6});
    assert.Equal(t, []int {2, 4, 1}, result);
}

func Test_GenomicRangeQuery_Sample2(t *testing.T) {
    result := Solution("A", []int{0}, []int {0});
    assert.Equal(t, []int {1}, result);
}

//   "CAGCCTA"
//    2132241
//      [ ]
// A: 0111112
// C: 1112333
// G: 0011111
// T: 0000011
// [2-4]
//    0011100