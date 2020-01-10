package luhn

func Valid(s string) bool {

	if len(s) < 2 {
		return false
	}
	var sumDigits, l int
	for i := len(s) - 1; i >= 0; i-- {
		d := s[i]
		switch {
		case d == ' ':
			continue
		case d >= '0' && d <= '9':
			l++
			n := int(d - '0')
			if l%2 == 0 {
				n = 2 * n
				if n > 9 {
					n -= 9
				}

			}
			sumDigits += n
		default:
			return false
		}
	}
	return l > 1 && sumDigits%10 == 0

}
