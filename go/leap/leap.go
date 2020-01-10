package leap

// IsLeapYear computes whether a given year is a leap year
func IsLeapYear(y int) bool {

	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}
