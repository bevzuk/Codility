package Ex04CountingElementsMissingInteger

// https://codility.com/demo/results/training7UGNT9-W25/

func Solution(A []int) int {
    B := make([]int, len(A));
    for _, a := range A {
        if (a-1 < len(A) && a > 0) {
            B[a-1] = 1;
        }
    }
    
    for i, b := range B {
        if (b == 0) {
            return i+1;
        }
    }
    return len(A) + 1;
}
