package ids

import "errors"

var (
	ErrInvalidID = errors.New("invalid id")
	ErrNilID     = errors.New("id is nil")
	ErrWrongSize = errors.New("id must be 8 bytes")
)
