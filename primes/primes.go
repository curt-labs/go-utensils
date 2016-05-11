package primes

func Primes(start int, end int) ([]int, []int) {
	var primes []int
	var nonprimes []int
	for i := start; i <= end; i++ {
		var done bool
		for j := 2; j < i; j++ {
			if i%j == 0 {
				nonprimes = append(nonprimes, i)
				done = true
				break
			}
		}
		if !done {
			primes = append(primes, i)
		}
	}
	return primes, nonprimes
}
