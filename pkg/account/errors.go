package account

type DatabaseError struct {
	Message string
}

func (e DatabaseError) Error() string {
	return e.Message
}

type AccountNotFoundError struct {
	Email string
}

func (e AccountNotFoundError) Error() string {
	return e.Email
}
