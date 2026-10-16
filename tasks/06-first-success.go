/*
ЗАДАЧА 7. ПОЛУЧИТЬ ПЕРВЫЙ УСПЕШНЫЙ ОТВЕТ

Что необходимо доработать:
  - Запрашивать файл во всех хранилищах одновременно.
  - Вернуть первый успешный ответ, остановив остальные запросы.
  - Ошибки хранилищ игнорировать: остальные запросы продолжают работу.
  - Если все запросы завершились без успешного ответа, вернуть ErrNotFound.
  - Если за одну секунду успешного ответа нет, остановить запросы
  - При отмене входного контекста остановить запросы и вернуть ctx.Err().
*/
package tasks

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("файл не найден")

type StorageFile struct {
	Data []byte
}

type Storage interface {
	Get(ctx context.Context, path string) (StorageFile, error)
}

func GetFile(ctx context.Context, path string, storages []Storage) (StorageFile, error) {
	for _, storage := range storages {
		file, err := storage.Get(ctx, path)
		if err == nil {
			return file, nil
		}
	}
	return StorageFile{}, ErrNotFound
}
