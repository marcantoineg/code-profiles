package utils

import "os"

// package utils exposes various useful functions. well, function in that case...

// Check panics if error is not nil.
func Check(e error) {
	if e != nil {
		println(e.Error())
		os.Exit(0)
	}
}
