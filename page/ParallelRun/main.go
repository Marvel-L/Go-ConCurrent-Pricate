package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func ParallelRun[T any, U any](tasks []T, numG int, fn func(context.Context, T) (U, error)) ([]U, error) {
	result := make([]U, len(tasks))
	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(numG)

	for i := range tasks {
		index := i
		g.Go(func() (err error) {
			defer func() {
				if panicErr := recover(); panicErr != nil {
					err = fmt.Errorf("panic: %v", panicErr)
				}
			}()
			res, err := fn(ctx, tasks[index])
			if err != nil {
				return err
			}
			result[index] = res
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return result, nil
}
