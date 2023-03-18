package utils

// package utils exposes various useful functions. well, function in that case...

// Check panics if error is not nil.
func Check(e error) {
	if e != nil {
		panic(e.Error())
	}
}
