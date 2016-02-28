package Ex03TimeComplexityFrogJmp

import (
//    "math"
)

// https://codility.com/demo/results/trainingWWJEAF-TPA/

func Solution(X int, Y int, D int) int {
    if (D == 0) { 
        return 0;
    }
    
    jumps := (Y - X) / D;
    if (X + jumps * D < Y) {
        jumps++;
    }
    return jumps;
}
