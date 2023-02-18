package registration

type EmailFormatError struct {
	Email string
}

func (e *EmailFormatError) Error() string {
	return e.Email
}

type PasswordFormatError struct {
}

func (e *PasswordFormatError) Error() string {
	return "*****"
}

type EmailAlreadyExistsError struct {
	Email string
}

func (e *EmailAlreadyExistsError) Error() string {
	return e.Email
}
