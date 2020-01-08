package hamming

import "errors"

func Distance(a, b string) (int, error) {
	// if len(a) == 0 || len(b) == 0 {
	// 	return 0, errors.New(("Distance: no empty sequences allowed"))
	// }
	if len(a) != len(b) {
		return 0, errors.New(("Distance: size mismatch"))
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance += 1
		}
	}
	return distance, nil
}
