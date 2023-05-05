package stores

import "database/sql"

type Store struct {
	UserStore *UserStore
	AuthStore *AuthStore
}

func NewStore(conn *sql.DB) *Store {
	return &Store{
		NewUserStore(conn),
		NewAuthStore(conn),
	}
}
