package account

import "time"

type Account struct {
	UserID       string
	Email        string
	Password     []byte
	Salt         []byte
	CreationDate time.Time
	UpdateDate   time.Time
}
