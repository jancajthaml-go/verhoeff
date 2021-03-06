// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Verhoeff_algorithm
//
package verhoeff

import "fmt"

var (
	mult = [10][10]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {1, 2, 3, 4, 0, 6, 7, 8, 9, 5}, {2, 3, 4, 0, 1, 7, 8, 9, 5, 6}, {3, 4, 0, 1, 2, 8, 9, 5, 6, 7}, {4, 0, 1, 2, 3, 9, 5, 6, 7, 8}, {5, 9, 8, 7, 6, 0, 4, 3, 2, 1}, {6, 5, 9, 8, 7, 1, 0, 4, 3, 2}, {7, 6, 5, 9, 8, 2, 1, 0, 4, 3}, {8, 7, 6, 5, 9, 3, 2, 1, 0, 4}, {9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}
	perm = [10][10]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {1, 5, 7, 6, 2, 8, 3, 0, 9, 4}, {5, 8, 0, 3, 7, 9, 6, 1, 4, 2}, {8, 9, 1, 6, 0, 4, 3, 5, 2, 7}, {9, 4, 5, 3, 1, 2, 6, 8, 7, 0}, {4, 2, 8, 6, 5, 7, 3, 9, 0, 1}, {2, 7, 9, 3, 8, 0, 6, 4, 1, 5}, {7, 0, 4, 6, 9, 1, 3, 2, 5, 8}}
	inv  = [10]int{0, 4, 3, 2, 1, 5, 6, 7, 8, 9}
	zero = int('0')
)

// Digit returns verhoeff digit for given numeric string
func Digit(cc string) (int, error) {
	var (
		d int
		i int
		x int
		l int = len(cc)
	)
scan:
	d = int(cc[l-i-1]) - zero
	if d < 0 || d > 9 {
		return 1, fmt.Errorf("string must contain only digits")
	}
	i++
	x = mult[x][perm[(i - (i/8)*8)][d]]
	if i != l {
		goto scan
	}
	return inv[x], nil
}

// Validate returns true if numeric string is signed with valid luhn digit
func Validate(cc string) bool {
	var (
		d int
		i int
		x int
		l int = len(cc)
	)
scan:
	d = int(cc[l-i-1]) - zero
	if d < 0 || d > 9 {
		return false
	}
	x = mult[x][perm[(i - (i/8)*8)][d]]
	i++
	if i != l {
		goto scan
	}
	return x == 0
}

// Generate signs string with verhoeff digit
func Generate(cc string) (string, error) {
	digit, err := Digit(cc)
	if err != nil {
		return cc, err
	}
	return cc + string(rune(digit+48)), nil
}
