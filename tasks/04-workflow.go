/*
ЗАДАЧА 5. ОБРАБАТЫВАТЬ РАЗНЫЕ ОПЕРАЦИИ ПАРАЛЛЕЛЬНО

Что необходимо доработать:
  - Разные операции должны выполняться параллельно.
  - Команды одной операции должны выполняться по очереди, в порядке получения.
  - При закрытии канала закончить обработку всех полученных команд.
  - При отмене контекста остановить работу. Необработанные команды можно отбросить.
  - Ошибки не обрабатывать.
*/
package tasks

import "context"

type WorkflowFile interface {
	Delete(ctx context.Context)
	Add(ctx context.Context)
	Update(ctx context.Context)
}

type Action int

const (
	Delete Action = iota
	Add
	Update
)

type Command struct {
	Action       Action
	WorkflowFile WorkflowFile
}

func RunWorkflow(ctx context.Context, input <-chan Command) {
	for command := range input {
		switch command.Action {
		case Delete:
			command.WorkflowFile.Delete(ctx)
		case Add:
			command.WorkflowFile.Add(ctx)
		case Update:
			command.WorkflowFile.Update(ctx)
		}
	}
}
