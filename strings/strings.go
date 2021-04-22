// A simple "strings" package to demonstrate Go's package
package strings

func internalFunction() {
	// In Go, a name is exported if it
	// begins with a capital letter
}

// Must begin with a capital letter in
// order to be exported
func CountOddEven(s string) (odds, evens int) {
	odds, evens = 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return
}
