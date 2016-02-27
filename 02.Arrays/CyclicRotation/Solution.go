package Ex02ArraysCyclicRotation

import (
)

// https://codility.com/demo/results/trainingKXJ255-JD6/

func Solution(A []int, K int) []int {
    len := len(A);
    B := make([]int, len);
    for index, value := range A {
        B[(index + K) % len] = value;
    }
    return B;
}