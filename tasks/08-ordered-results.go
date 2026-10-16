/*
ЗАДАЧА 9. ОБРАБОТАТЬ ФАЙЛЫ ПАРАЛЛЕЛЬНО И СОХРАНИТЬ ПОРЯДОК РЕЗУЛЬТАТОВ

Что необходимо доработать:
  - Каждый файл обрабатывать в отдельной горутине.
  - Получить результаты из Process и вернуть их в порядке исходного списка.
  - Ошибки не обрабатывать.
*/
package tasks

import "context"

type Result struct {
	Path string
}

type OrderedFile interface {
	Process(ctx context.Context) Result
}

func ProcessFilesOrdered(ctx context.Context, files []OrderedFile) []Result {
	results := make([]Result, 0, len(files))
	for _, file := range files {
		results = append(results, file.Process(ctx))
	}
	return results
}
