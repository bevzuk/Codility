package Ex05PrefixSumsPassingCars

// https://codility.com/demo/results/trainingG2HE3C-G7E/

func Solution(A []int) int {
    eastCars := 0;
    passingCarsCount := 0;
    
    for _, a := range A {
        if (a == 0) {
            eastCars++;
        } else {
            passingCarsCount += eastCars;
        }
        
        if (passingCarsCount > 1000000000) {
            return -1;
        }
    }
    
    return passingCarsCount;
}