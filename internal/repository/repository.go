package repository

import (
	"context"
	"github.com/y3g0r/rehearsals-api-go/internal/domain"
	"sync"
)

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		Users: make([]domain.User, 0),
	}
}

type InMemoryUserRepository struct {
	Users []domain.User
	Lock  sync.Mutex
}

func (r *InMemoryUserRepository) CreateUser(ctx context.Context, u domain.User) error {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	for _, user := range r.Users {
		if user.ID() == u.ID() {
			return ErrUserIDCollision
		}
	}
	r.Users = append(r.Users, u)
	return nil
}

func (r *InMemoryUserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	for _, user := range r.Users {
		if user.Email() == email {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (r *InMemoryUserRepository) getUserIdxById(ctx context.Context, id string) (int, error) {
	for i, user := range r.Users {
		if user.ID() == id {
			return i, nil
		}
	}
	return 0, ErrUserNotFound
}

func (r *InMemoryUserRepository) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	idx, err := r.getUserIdxById(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.Users[idx], nil
}

func (r *InMemoryUserRepository) UpdateUser(ctx context.Context, u domain.User) error {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	idx, err := r.getUserIdxById(ctx, u.ID())
	if err != nil {
		return err
	}

	r.Users[idx] = u
	return nil

}
