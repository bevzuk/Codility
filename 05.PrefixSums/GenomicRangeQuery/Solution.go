package Ex03PrefixSumsGenomicRangeQuery;

import (
)

// https://codility.com/demo/results/trainingHGMZN6-TDH/

func Solution(S string, P []int, Q []int) []int {
    impactsCounts := make([][]int, 4);
    impactsCounts[0] = make([]int, len(S)+1);
    impactsCounts[1] = make([]int, len(S)+1);
    impactsCounts[2] = make([]int, len(S)+1);
    impactsCounts[3] = make([]int, len(S)+1);
    
    for i, s := range S {
        impactsCounts[0][i+1] = impactsCounts[0][i];
        impactsCounts[1][i+1] = impactsCounts[1][i];
        impactsCounts[2][i+1] = impactsCounts[2][i];
        impactsCounts[3][i+1] = impactsCounts[3][i];
        switch s {
            case 'A': impactsCounts[0][i+1]++;
            case 'C': impactsCounts[1][i+1]++;
            case 'G': impactsCounts[2][i+1]++;
            case 'T': impactsCounts[3][i+1]++;
        }
    }
    
    mins := make([]int, len(P));
    for i, p := range P {
        q := Q[i]+1;
        
        if (impactsCounts[0][p] < impactsCounts[0][q]) { 
            mins[i] = 1;
            continue; 
        }
        if (impactsCounts[1][p] < impactsCounts[1][q]) { 
            mins[i] = 2;
            continue; 
        }
        if (impactsCounts[2][p] < impactsCounts[2][q]) { 
            mins[i] = 3;
            continue;
        }
        mins[i] = 4;
    }
    
    return mins;
}
