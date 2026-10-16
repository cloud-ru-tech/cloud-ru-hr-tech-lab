/*
ЗАДАЧА 12. ЗАКОНЧИТЬ ПРИНЯТЫЕ ЗАДАНИЯ ПЕРЕД ОСТАНОВКОЙ

Что необходимо доработать:
  - Сейчас после сигнала остановки процессы не завершаються
*/
package tasks

import "context"

type GracefulJob interface {
	Process(ctx context.Context)
}

func Run(ctx context.Context, input <-chan GracefulJob, stop <-chan struct{}, workers int) {
	for {
		select {
		case <-stop:
			return
		case job, ok := <-input:
			if !ok {
				return
			}
			go job.Process(ctx)
		}
	}
}
