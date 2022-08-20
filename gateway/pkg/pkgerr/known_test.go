package pkgerr_test

import (
	"fmt"
	"testing"

	"github.com/isutare412/meetup/gateway/pkg/pkgerr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fmtOriginErr = "complex error ocurred under the hood"
	fmtSimpleErr = "something(id=%d) got wrong"
)

func TestKnown(t *testing.T) {
	const (
		id = 412
	)

	err := funcOne(id)
	kerr := pkgerr.AsKnown(err)

	require.NotNil(t, kerr)
	assert.Equal(t, fmt.Sprintf(fmtSimpleErr, id), kerr.SimpleError())
	assert.Equal(t, fmtOriginErr, kerr.Error())
}

func funcOne(id int) error {
	err := funcTwo(id)
	return fmt.Errorf("from funcTwo: %w", err)
}

func funcTwo(id int) error {
	err := originError(id)
	return fmt.Errorf("from originError: %w", err)
}

func originError(id int) error {
	err := fmt.Errorf(fmtOriginErr)
	return pkgerr.Known{
		Origin: err,
		Simple: fmt.Errorf(fmtSimpleErr, id),
	}
}
