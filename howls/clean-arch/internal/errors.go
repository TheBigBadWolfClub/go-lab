package internal

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrToolUnavailable   = errors.New("tool unavailable")
	ErrNotImplemented    = errors.New("use Case Not Implemented")
	ErrInvalidEntity     = errors.New("invalid entity data")
	ErrInconsistentState = errors.New("unexpected entity state")
	ErrFailPayment       = errors.New("unable to process payment")
	ErrMaxAllowedReached = errors.New("customer reached max assignments allowed")
)
