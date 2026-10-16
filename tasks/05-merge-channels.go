/*
ЗАДАЧА 6. ОБЪЕДИНИТЬ НЕСКОЛЬКО КАНАЛОВ В ОДИН

Что необходимо доработать:
  - Читать входные каналы параллельно и передавать данные в выходной канал.
  - Ошибки не обрабатывать.
*/
package tasks

import "context"

type MergedFile struct {
	Path string
}

func MergeChannels(ctx context.Context, inputs []<-chan MergedFile) <-chan MergedFile {
	output := make(chan MergedFile)

	for _, input := range inputs {
		for file := range input {
			output <- file
		}
	}

	return output
}
