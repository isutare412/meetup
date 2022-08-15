package perror

import "fmt"

type Known struct {
	Errno  Errno
	Source error
}

func (err Known) Error() string {
	if err.Source == nil {
		return fmt.Sprintf("errno: %d", err.Errno)
	}
	return err.Source.Error()
}
