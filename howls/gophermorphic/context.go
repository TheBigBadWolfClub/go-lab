package gophermorphic

import (
	contextpkg "context"
	"time"
)

/*
* Create a function that simulates a long-running operation.
* Use context cancellation to stop the operation early.
 */

const (
	canceled  = "Canceled"
	completed = "Operation completed"
)

// longRunningOperation simulates a long-running operation.
func longRunningOperation(ctx contextpkg.Context, duration time.Duration, resultChan chan string) {
	select {
	case <-ctx.Done():
		resultChan <- canceled
		return
	case <-time.After(duration):
		resultChan <- completed
	}
}
