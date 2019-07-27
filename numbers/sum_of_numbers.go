package numbers

func sumOfNumbers(n int) int {
	if n < 0 {
		return sumOfNumbers(-n)
	}
	if n < 10 {
		return n
	}
	return n%10 + sumOfNumbers(n/10)
}
