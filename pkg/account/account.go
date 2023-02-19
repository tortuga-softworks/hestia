package account

import "time"

type Account struct {
	UserID       string
	Email        string
	PasswordHash []byte
	CreationDate time.Time
	UpdateDate   time.Time
}
