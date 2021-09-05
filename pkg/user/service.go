package user

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	// GetUser(ctx context.Context, id string) (*User, error)
	// UpdateUser(ctx context.Context, id string, user *User) (*User, error)
	// ArchiveUser(ctx context.Context, id string) (*User, error)
}

type UserService interface {
	CreateUser(user *User) (*User, error)
}

type service struct {
	ur UserRepository
}

func NewUserService(ur UserRepository) UserService {
	return &service{ur: ur}
}

func (s *service) CreateUser(user *User) (*User, error) {
	return s.ur.CreateUser(context.TODO(), user)
}
