package armstrong

// Binary exponentation, i want to avoid the float conversion for the golang math.Pow function
func exp(b, e int) int {
	result := 1
	for e > 0 {
		if (e & 1) == 1 {
			result *= b
		}
		b *= b
		e >>= 1
	}
	return result
}
func IsNumber(n int) bool {
	var numDigits, sum int
	m := n
	for m > 0 {
		numDigits++
		m /= 10
	}
	m = n
	for m > 0 {
		d := m % 10
		m /= 10
		sum += exp(d, numDigits)
	}
	return sum == n
}
