package pkgerr

type Errno int32

const (
	ErrnoNone Errno = iota
	ErrnoEntityNotFound
)
