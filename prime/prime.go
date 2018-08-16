package prime

import "math"

func GetPrimeNumbers(n int) []int {
	var results []int
	for i := 0; len(results) < n; i++ {
		if IsPrime(i) {
			results = append(results, i)
		}
	}
	return results
}

func IsPrime(val int) bool {
	sqrRoot := int(math.Floor(math.Sqrt(float64(val))))
	for i := 2; i <= sqrRoot; i++ {
		if val%i == 0 {
			return false
		}
	}
	return val > 1
}
