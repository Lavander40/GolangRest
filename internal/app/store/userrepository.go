package store

import (
	"github.com/lavander40/golang_rest/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.BeforeCreate(); err != nil {
		return nil, err
	}

	result, err := r.store.db.Exec("INSERT INTO users (email, encryptedPassword) VALUES (?, ?)",
		user.Email,
		user.EncryptedPassword,
	)
	if err != nil {
		return nil, err
	}

	LastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.Id = int(LastId)
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, encryptedPassword FROM users WHERE email = ?",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}
	return u, nil
}
