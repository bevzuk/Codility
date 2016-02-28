package Ex04CountingElementsPermCheck

// https://codility.com/demo/results/training4X8BPG-AWM/

func Solution(A []int) int {
    B := make([]int, len(A));
    
    for _, a := range A {
        a1 := a - 1;
        if (a1 >= 0 && a1 < len(A)) {
            B[a1] = 1;
        }
    }
    
    for _, b := range B {
        if (b == 0) {
            return 0;
        }
    }
    
    return 1;
}
