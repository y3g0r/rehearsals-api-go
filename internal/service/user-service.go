package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/y3g0r/rehearsals-api-go/internal/domain"
)

type userRepository interface {
	CreateUser(ctx context.Context, u domain.User) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	//GetUserByID(ctx context.Context, id string) (domain.User, error)
	//UpdateUser(ctx context.Context, u domain.User) error
}

type cryptoService interface {
	HashPassword(password string) (string, error)
	ComparePasswordAndHash(password string, hashedPassword string) error
}

type UserService struct {
	usersRepo userRepository
	cryptoSvc cryptoService
}

func (s *UserService) Authenticate(ctx context.Context, p AuthenticationParams) (domain.User, error) {
	user, err := s.usersRepo.GetUserByEmail(ctx, p.Username)
	if err != nil {
		return nil, err
	}

	if err = s.comparePasswordAndHash(p.Password, user.HashedPassword()); err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserService(usersRepo userRepository, cryptoSvc cryptoService) *UserService {
	return &UserService{
		usersRepo: usersRepo,
		cryptoSvc: cryptoSvc,
	}
}

func (s *UserService) CreateUser(ctx context.Context, p CreateUserParams) error {
	hashedPassword, err := s.hashPassword(p.Password)
	if err != nil {
		return fmt.Errorf("couldn't hash password: %w", err)
	}

	for retries := 3; retries > 0; retries-- {
		id := s.generateUserId()
		user := domain.NewSimpleUser(id, hashedPassword)

		if err = s.usersRepo.CreateUser(ctx, user); err == nil {
			return nil
		}
	}

	return fmt.Errorf("couldn't create user: %w", err)
}

func (s *UserService) hashPassword(password string) (string, error) {
	return s.cryptoSvc.HashPassword(password)
}

func (s *UserService) generateUserId() string {
	// generate uuid and return as string
	return uuid.New().String()
}

func (s *UserService) comparePasswordAndHash(password string, hashedPassword string) error {
	if err := s.cryptoSvc.ComparePasswordAndHash(password, hashedPassword); err != nil {
		return fmt.Errorf("passwords don't match: %w", err)
	}
	return nil
}
