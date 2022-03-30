package finder_test

import (
	"errors"
	"testing"

	"github.com/dmytro-vovk-f3/procid/internal/finder"
	"github.com/stretchr/testify/assert"
)

func TestResolve(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func()
		teardown func()
		input    string
		expect   string
		err      error
	}{
		{
			name:   "Empty path",
			input:  "",
			expect: "",
			err:    finder.ErrEmptyPath,
		},
		{
			name:   "Existing dir",
			input:  "/tmp",
			expect: "",
			err:    errors.New("/tmp points to a directory"),
		},
		{
			name:   "Existing binary",
			input:  "/bin/sh",
			expect: "/usr/bin/dash",
		},
	}

	for _, tt := range testCases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()

				if tt.teardown != nil {
					t.Cleanup(tt.teardown)
				}
			}

			result, err := finder.Resolve(tt.input)
			if assert.Equal(t, tt.err, err) {
				assert.Equal(t, tt.expect, result)
			}
		})
	}
}
