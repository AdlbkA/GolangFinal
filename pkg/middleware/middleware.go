package middleware

import (
	"context"
	"errors"
	"log"
)

func WaitContextCancel(ctx context.Context, method string, fn func() error) error {
	errCh := make(chan error, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("%v panic recovered: %v", method, r)
				errCh <- errors.New("panic recovered")
			}
		}()

		errCh <- fn()
	}()

	select {
	case <-ctx.Done():
		return errors.New("reached timeout")
	case err := <-errCh:
		return err
	}
}
