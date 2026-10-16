/*
ЗАДАЧА 8. ОБРАБАТЫВАТЬ НАКОПЛЕННЫЕ ФАЙЛЫ ПО ТАЙМЕРУ

Что необходимо доработать:
  - Не ждать все из канала до закрытия а каждые 2 секунды вызвать process на то что уже пришло
  - С ошибками не работаем
*/
package tasks

import (
	"context"
	"time"
)

type PeriodicFile struct {
	Path string
}

func processPeriodicBatch(ctx context.Context, files []PeriodicFile) {
	// ...
}

func ProcessPeriodically(ctx context.Context, input <-chan PeriodicFile, interval time.Duration) {
	var files []PeriodicFile
	for file := range input {
		files = append(files, file)
	}
	processPeriodicBatch(ctx, files)
}
