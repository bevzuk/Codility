package Ex03TimeComplexityPermMissingElem

// https://codility.com/demo/results/trainingD9QTXA-BE4/

func Solution(A []int) int {
    sum := 0;
    for _, a := range A {
        sum += a;
    }

    expectedSum := 0;
    for i:=1; i<=len(A)+1; i++ {
        expectedSum += i;
    }

    return expectedSum - sum;
}
