package gophermorphic

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFinished(t *testing.T) {
	t.Run("TestWithTimeout", func(t *testing.T) {
		resultChan := make(chan string)
		go longRunningOperation(context.Background(), 1*time.Second, resultChan)
		assert.Equal(t, completed, <-resultChan)
	})
}

func TestWithCancel(t *testing.T) {
	t.Run("TestWithCancel", func(t *testing.T) {
		resultChan := make(chan string)
		ctx, cancelFunc := context.WithCancel(context.Background())
		go longRunningOperation(ctx, 5*time.Second, resultChan)
		go func() {
			time.Sleep(1 * time.Second)
			cancelFunc()
		}()
		assert.Equal(t, canceled, <-resultChan)
	})
}

func TestWithCancelCause(t *testing.T) {
	t.Run("TestWithCancelCause", func(t *testing.T) {
		resultChan := make(chan string)
		ctx, cancelCauseFunc := context.WithCancelCause(context.Background())
		go longRunningOperation(ctx, 5*time.Second, resultChan)
		go func() {
			time.Sleep(1 * time.Second)
			cancelCauseFunc(errors.New("canceled"))
		}()
		assert.Equal(t, canceled, <-resultChan)
		assert.Equal(t, context.Cause(ctx), errors.New("canceled"))
	})
}

func TestWithDeadline(t *testing.T) {
	t.Run("TestWithDeadline", func(t *testing.T) {
		resultChan := make(chan string)
		ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
		defer cancelFunc()
		go longRunningOperation(ctx, 5*time.Second, resultChan)
		assert.Equal(t, canceled, <-resultChan)
	})
}

func TestWithDeadlineCause(t *testing.T) {
	t.Run("TestWithDeadlineCause", func(t *testing.T) {
		resultChan := make(chan string)
		ctx, cancelFunc := context.WithDeadlineCause(context.Background(), time.Now().Add(1*time.Second), errors.New("deadline exceeded"))
		defer cancelFunc()
		go longRunningOperation(ctx, 5*time.Second, resultChan)

		assert.Equal(t, canceled, <-resultChan)
		assert.Equal(t, context.Cause(ctx), errors.New("deadline exceeded"))
	})
}

func TestWithTimeout(t *testing.T) {
	t.Run("TestWithTimeout", func(t *testing.T) {
		resultChan := make(chan string)
		ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancelFunc()
		go longRunningOperation(ctx, 5*time.Second, resultChan)
		assert.Equal(t, canceled, <-resultChan)
	})
}

func TestWithTimeoutCause(t *testing.T) {
	t.Run("TestWithTimeoutCause", func(t *testing.T) {
		resultChan := make(chan string)
		ctx, cancelFunc := context.WithTimeoutCause(context.Background(), 1*time.Second, errors.New("timeout"))
		defer cancelFunc()
		go longRunningOperation(ctx, 5*time.Second, resultChan)

		assert.Equal(t, canceled, <-resultChan)
		assert.Equal(t, context.Cause(ctx), errors.New("timeout"))
	})
}

func TestMultipleCancel(t *testing.T) {
	t.Run("TestMultipleCancel", func(t *testing.T) {
		resultChan := make(chan string, 10)
		ctx, cancelFunc := context.WithCancel(context.Background())

		for i := 0; i < 10; i++ {
			go longRunningOperation(ctx, 5*time.Second, resultChan)
		}

		go func() {
			time.Sleep(1 * time.Second)
			cancelFunc()
		}()

		for i := 0; i < 10; i++ {
			assert.Equal(t, canceled, <-resultChan)
		}
	})
}
