package diffsquares

func Difference(n int) int {
	if n<2 {
		return 0
	}
	return int(n*(n+1)*(n-1)*(3*n+2)/12)
}

func SquareOfSum(n int) (sum int) {
	sum = n*n*(n+1)*(n+1) / 4
	return 
}

func SumOfSquares(n int) (sum int) {
	sum = (2*n+1)*n*(n+1) / 6
	return
}