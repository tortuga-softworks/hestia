package registration

import (
	"context"
	"errors"
	"net/mail"

	"github.com/tortuga-softworks/hestia/pkg/account"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	accountStore account.AccountStore
}

func NewRegistrationService(accountStore account.AccountStore) (*RegistrationService, error) {
	if accountStore == nil {
		return nil, errors.New("could not create a registration service: account store is nil")
	}

	return &RegistrationService{accountStore}, nil
}

func (rs *RegistrationService) SignUp(ctx context.Context, email, password string) (string, error) {

	if !verifyEmailFormat(email) {
		return "", &EmailFormatError{email}
	}

	if !verifyPasswordFormat(password) {
		return "", &PasswordFormatError{}
	}

	_, err := rs.accountStore.FindByEmail(ctx, email)
	if err != nil {
		switch err.(type) {
		case *account.AccountNotFoundError:
			pwd, salt, err := hashPassword(password)

			if err != nil {
				return "", errors.New("could not secure password")
			}

			account := account.Account{
				Email:    email,
				Password: pwd,
				Salt:     salt,
			}

			err = rs.accountStore.Create(ctx, &account)

			if err != nil {
				return "", err
			}

			createdAccount, err := rs.accountStore.FindByEmail(ctx, email)

			if err != nil {
				return "", err
			}

			return createdAccount.UserID, nil
		default:
			return "", err
		}
	} else {
		return "", &EmailAlreadyExistsError{email}
	}
}

func verifyEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func verifyPasswordFormat(password string) bool {
	return len(password) >= 6
}

func hashPassword(password string) ([]byte, []byte, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}
	return hashedPassword, salt, nil
}
