package mock

import (
	"context"

	"github.com/mattismoel/konnekt/internal/storage"
)

var users = []storage.User{
	{ID: 1, FirstName: "John", LastName: "Doe", Email: "test@mail.com", PasswordHash: []byte("Hash123")},
	{ID: 2, FirstName: "Willem", LastName: "Dafoe", Email: "wilfoe@mail.com", PasswordHash: []byte("Hash123")},
	{ID: 3, FirstName: "Max", LastName: "Maxwell", Email: "maxwell@mail.com", PasswordHash: []byte("Hash123")},
}

func (s mockStorage) InsertUser(ctx context.Context, user storage.User) (storage.User, error) {
	users = append(users, user)

	user.ID = int64(len(users))

	return user, nil
}

func (s mockStorage) DeleteUser(ctx context.Context, id int64) error {
	for i, u := range users {
		if u.ID == id {
			users = removeFromSlice(users, i)
		}
	}

	return nil
}

func (s mockStorage) UpdateUser(ctx context.Context, id int64, newUser storage.User) (storage.User, error) {
	for i, u := range users {
		if u.ID == id {
			users[i] = storage.User{
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Email:     u.Email,
			}

			return users[i], nil
		}
	}

	return storage.User{}, storage.ErrNotFound
}
