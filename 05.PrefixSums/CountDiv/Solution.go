package Ex05PrefixSumsCountDiv

// https://codility.com/demo/results/trainingY3MMFH-VA7/

func Solution(A int, B int, K int) int {
    offset := A % K;
    if (offset == 0) {
        A = A - K;
    } else {
        A = A - offset;
    }
    
    return (B - A) / K;
}