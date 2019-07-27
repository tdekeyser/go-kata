package numbers

func SumOfNumbers(n int) int {
	if n < 0 {
		return SumOfNumbers(-n)
	}
	if n < 10 {
		return n
	}
	return n%10 + SumOfNumbers(n/10)
}
