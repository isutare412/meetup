package pkgerr

type Errno int32

const (
	ErrnoInternal Errno = iota
	ErrnoEntityNotFound
)
