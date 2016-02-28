package Ex04CountingElementsFrogRiverOne

// https://codility.com/demo/results/training7Y5T8Y-JUF/

func Solution(X int, A []int) int {
    leafs := make([]int, X);
    leafsLeft := X;
    for i, a := range A {
        if (leafs[a-1] == 0) {
            leafs[a-1] = 1;
            leafsLeft--;
        }
        if leafsLeft == 0 {
            return i;
        }
    }
    return -1;
}
