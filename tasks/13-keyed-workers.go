/*
ЗАДАЧА 14. ОБРАБАТЫВАТЬ ОДИН ФАЙЛ ПО ОЧЕРЕДИ, РАЗНЫЕ — ПАРАЛЛЕЛЬНО

Что необходимо доработать:
  - Сейчас все файлы обрабатываются по очереди. Нужно использовать workers обработчиков.
  - Задания одного файла должны выполняться в порядке получения,
    а задания разных файлов могут выполняться параллельно.
*/
package tasks

import "context"

type KeyedFile interface {
	PathHash() uint64
	Process(ctx context.Context)
}

func ProcessFilesByKey(ctx context.Context, input <-chan KeyedFile, workers int) {
	for file := range input {
		file.Process(ctx)
	}
}
