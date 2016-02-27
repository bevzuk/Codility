package Ex02ArraysOddOccurrencesInArray

// https://codility.com/demo/results/trainingKR7WXP-DN7/

func Solution(A []int) int {
    result := 0;
    for _, x := range A {
        result ^= x;
    }
    return result;
}
