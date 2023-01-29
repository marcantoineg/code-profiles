package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Check(t *testing.T) {
	testRuns := []struct {
		testName string
		err      error

		expectsPanic bool
	}{
		{"Check with nil as error should not panic", nil, false},
		{"Check with actual error should panic", errors.New("some error"), true},
	}

	for _, tr := range testRuns {
		t.Run(tr.testName, func(t *testing.T) {
			if tr.expectsPanic {
				assert.Panics(t, func() { Check(tr.err) })
			} else {
				Check(tr.err)
				assert.Nil(t, tr.err)
			}
		})
	}
}
