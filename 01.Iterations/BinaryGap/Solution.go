package Ex01IterationsBinaryGap;

import (
    "strconv"
)

// https://codility.com/demo/results/trainingKF78KJ-Z89/

// Solution finds a binary gap
func Solution(N int) int {
    binary := strconv.FormatInt(int64(N), 2);
    return countZeros(binary, findFirst1(binary), findLast1(binary));
}

func countZeros(s string, from int, to int) int {
    max := 0;
    count := 0;
    for i := from; i <= to; i++ {
        if s[i] == '0' {
            count++;
        } else {
            if (count > max) {
                max = count;
            }
            count = 0;
        }
    } 
    return max;
}

func findFirst1(s string) int {
    index := 0;
    for index < len(s) && s[index] == '0' { 
        index++;
    }
    return index;
}

func findLast1(s string) int {
    index := len(s) - 1;
    for index >= 0 && s[index] == '0' { 
        index--;
    }
    return index;
}
