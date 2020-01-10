package grains

import "errors"

// Square computes the number of grains on the chessboards (n must be between 1 and 64)
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("out of bounds, n must be between 1 and 64")
	}
	return 1 << (n - 1), nil
}

// Total computes the number of grains on the whole chessboard (1+2+4+....)
func Total() uint64 {

	return 1<<64 - 1
}
