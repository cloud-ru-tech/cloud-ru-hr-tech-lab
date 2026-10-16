/*
ЗАДАЧА 1. ОБРАБОТАТЬ ФАЙЛЫ НЕСКОЛЬКИМИ ВОРКЕРАМИ

Что необходимо доработать:
  - Обрабатывать файлы параллельно, используя пул из workers воркеров.
  - Обработать каждый элемент списка ровно один раз.
  - Функция должна дождаться завершения всей обработки.
  - Обработка ошибок не требуется.
*/
package tasks

import "context"

type WorkerFile interface {
	Process(ctx context.Context)
}

func ProcessFilesWithWorkers(ctx context.Context, files []WorkerFile, workers int) {
	for _, file := range files {
		file.Process(ctx)
	}
}
