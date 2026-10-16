/*
ЗАДАЧА 15. ОБЪЕДИНИТЬ ОДНОВРЕМЕННЫЕ ЗАПРОСЫ ОДНОГО ПОЛЬЗОВАТЕЛЯ

Что необходимо доработать:
  - Сейчас каждый вызов GetUser обращается к БД. Если 100 горутин
    одновременно запрашивают одного пользователя, в БД уходят 100 запросов.
  - Нужно выполнять один запрос к БД для одного id, пока он выполняется.
  - Остальные горутины с тем же id должны дождаться этого запроса
    и получить того же пользователя или ту же ошибку.
  - Запросы разных пользователей должны выполняться независимо.
*/
package tasks

import "database/sql"

type User struct {
	ID   int
	Name string
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(id int) (User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name)
	return user, err
}
