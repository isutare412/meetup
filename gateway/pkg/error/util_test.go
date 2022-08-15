package perror

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryRollback(t *testing.T) {
	type testSet struct {
		name      string
		originErr error
		rollback  func() error
		expected  string
	}

	var testSets = []*testSet{
		{
			name:     "no-error",
			expected: "",
		},
		{
			name:      "rollback-success",
			originErr: fmt.Errorf("origin error"),
			rollback:  func() error { return nil },
			expected:  "origin error",
		},
		{
			name:      "rollback-failure",
			originErr: fmt.Errorf("origin error"),
			rollback:  func() error { return fmt.Errorf("rollback failed") },
			expected:  "rollback failed: origin error",
		},
	}

	for _, ts := range testSets {
		t.Run(ts.name, func(t *testing.T) {
			var got = TryRollback(ts.originErr, ts.rollback)
			if ts.expected == "" {
				assert.Nil(t, got)
			} else {
				assert.Equal(t, ts.expected, got.Error())
			}
		})
	}
}
