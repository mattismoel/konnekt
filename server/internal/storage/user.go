package storage

type User struct {
	ID           int64
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
}
