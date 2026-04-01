package collatzconjecture

import "errors"

func collatzConjuctureRec(n int, steps int) (int) {
    if (n == 1) {
        return steps
    } 
    if (n%2 == 0){
        return collatzConjuctureRec(n/2, steps+1)
    } 
    return collatzConjuctureRec(3*n+1, steps+1)
}

func CollatzConjecture(n int) (int, error) {
    if n<=0 {
        return 0, errors.New("Number is not a positive integer") 
    }
	return collatzConjuctureRec(n, 0), nil
}
