/*
ЗАДАЧА 2. ОБРАБОТАТЬ ОШИБКИ ДВУХ ОПЕРАЦИЙ

Сейчас ошибки игнорируются.

Что необходимо доработать:
  - Если одна операция вернула ошибку, другая должна остановиться.
  - Метод должен вернуть первую полученную ошибку.
*/
package tasks

import (
	"context"
	"sync"
)

func RunWorkers(ctx context.Context, first, second func(context.Context) error) error {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		first(ctx)
	}()

	go func() {
		defer wg.Done()
		second(ctx)
	}()

	wg.Wait()
	return nil
}
