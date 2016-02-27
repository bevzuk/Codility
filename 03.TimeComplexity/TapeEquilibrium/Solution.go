package Ex03TimeComplexityTapeEquilibrium

import (
    "math"
)

// https://codility.com/demo/results/trainingPZEKYG-S9P/

func Solution(A []int) int {
    left := 0;
    right := sum(A);
    min := math.MaxInt32;
    for i, a := range A {
        if (i == len(A) - 1) {
            continue;
        }
        
        left += a;
        right -= a;
        difference := abs(left - right);
        if (difference < min) {
            min = difference;
        }
    }
    return min;
}

func sum(A []int) int {
    sum := 0;
    for _, a := range A {
        sum += a;
    }
    return sum;
}

func abs(a int) int {
    if a >= 0 {
        return a;
    }
    return -a;
}