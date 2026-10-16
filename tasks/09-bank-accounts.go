/*
ЗАДАЧА 10. БЕЗОПАСНО РАБОТАТЬ СО СЧЕТАМИ КЛИЕНТОВ БАНКА

Что необходимо доработать:
  - Код хорошо работает когда работает пару человек но на проде хаос у всех на балансе
*/
package tasks

type Bank struct {
	balances map[string]int
}

func NewBank() *Bank {
	return &Bank{balances: make(map[string]int)}
}

func (b *Bank) Get(account string) int {
	return b.balances[account]
}

func (b *Bank) Add(account string, amount int) {
	b.balances[account] += amount
}

func (b *Bank) Transfer(from, to string, amount int) bool {
	b.Add(from, -amount)
	b.Add(to, amount)
	return true
}

func (b *Bank) TotalBalance() int {
	total := 0
	for _, balance := range b.balances {
		total += balance
	}
	return total
}
