package internal

import "errors"

var ErrNotFound = errors.New("not found")
var ErrToolUnavailable = errors.New("tool unavailable")
var ErrNotImplemented = errors.New("use Case Not Implemented")
var ErrInvalidEntity = errors.New("invalid entity data")
var ErrInconsistentState = errors.New("unexpected entity state")
var ErrFailPayment = errors.New("unable to process payment")
var ErrMaxAllowedReached = errors.New("customer reached max assignments allowed")
