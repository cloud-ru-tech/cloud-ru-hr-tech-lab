/*
ЗАДАЧА 4. УДАЛЯТЬ ТОЛЬКО ДУБЛИ ФАЙЛОВ

Сейчас метод удаляет каждый файл, который получает из канала.

Что необходимо доработать:
  - Первый файл с каждым хешем оставлять. Файлы с уже встречавшимся хешем удалять.
  - Читать файлы до закрытия канала или отмены контекста.
  - Ошибки не обрабатывать.

Пример: приходят A с хешем x, B с хешем y, C с хешем x.
A и B оставить, C удалить.
*/
package tasks

import "context"

type DuplicateFile interface {
	Hash() string
	Delete(ctx context.Context)
}

func RemoveDuplicates(ctx context.Context, input <-chan DuplicateFile) {
	for file := range input {
		file.Delete(ctx)
	}
}
