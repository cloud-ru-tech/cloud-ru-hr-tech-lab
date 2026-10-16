/*
ЗАДАЧА 11. СОЗДАТЬ ОБЩЕЕ ПОДКЛЮЧЕНИЕ ОДИН РАЗ

Что необходимо доработать:
  - Жалуются, что GetConnection работает очень долго.
    Почему мы вызываем Connect каждый раз, если можем вызвать его один раз?
*/
package tasks

import "context"

type Connection struct{}

type Connector interface {
	Connect() (*Connection, error)
}

type Manager struct {
	connector Connector
}

func NewManager(connector Connector) *Manager {
	return &Manager{connector: connector}
}

func (m *Manager) GetConnection(ctx context.Context) (*Connection, error) {
	return m.connector.Connect()
}
