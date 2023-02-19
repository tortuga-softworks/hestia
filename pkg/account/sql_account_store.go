package account

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type SqlAccountStore struct {
	db *sql.DB
}

func NewSqlAccountStore(db *sql.DB) (*SqlAccountStore, error) {
	if db == nil {
		return nil, errors.New("could not create an account store: database is nil")
	}

	return &SqlAccountStore{db: db}, nil
}

func (store *SqlAccountStore) FindByEmail(ctx context.Context, email string) (*Account, error) {
	row := store.db.QueryRowContext(ctx, "SELECT user_id, email, password, salt, creation_date, update_date FROM account WHERE email = $1", email)

	account := &Account{}
	err := row.Scan(&account.UserID, &account.Email, &account.Password, &account.Salt, &account.CreationDate, &account.UpdateDate)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &AccountNotFoundError{Email: email}
		}
		return nil, &DatabaseError{Message: err.Error()}
	}

	return account, nil
}

func (store *SqlAccountStore) Create(ctx context.Context, account *Account) (*Account, error) {
	query := "INSERT INTO account (email, password, salt) VALUES ($1, $2, $3) RETURNING user_id, creation_date, update_date"

	err := store.db.QueryRowContext(ctx, query, account.Email, account.Password, account.Salt).Scan(&account.UserID, &account.CreationDate, &account.UpdateDate)

	if err != nil {
		return nil, &DatabaseError{err.Error()}
	}

	return account, nil
}
