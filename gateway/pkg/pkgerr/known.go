package pkgerr

import (
	"errors"
	"fmt"
)

type Known struct {
	Errno  Errno
	Origin error
	Simple error
}

func (err Known) Error() string {
	if err.Origin == nil {
		if err.Simple == nil {
			return fmt.Sprintf("errno: %d", err.Errno)
		}
		return err.Simple.Error()
	}
	return err.Origin.Error()
}

func (err Known) SimpleError() string {
	if err.Simple == nil {
		if err.Origin == nil {
			return fmt.Sprintf("errno: %d", err.Errno)
		}
		return err.Origin.Error()
	}
	return err.Simple.Error()
}

func AsKnown(err error) *Known {
	var kerr Known
	if !errors.As(err, &kerr) {
		return nil
	}
	return &kerr
}
