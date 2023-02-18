package account

import "context"

type AccountStore interface {
	FindByEmail(ctx context.Context, email string) (*Account, error)
	Create(ctx context.Context, account *Account) error
}
