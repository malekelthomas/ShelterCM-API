package user

type UserRepository interface {
	Create(user *User) (*User, error)
	GetAll(role Role) ([]User, error)
	GetOneByID(id int64) (*User, error)
	// UpdateUser(ctx context.Context, id string, user *User) (*User, error)
	// ArchiveUser(ctx context.Context, id string) (*User, error)
}

type UserService interface {
	Create(user *User) (*User, error)
	GetAll(role Role) ([]User, error)
	GetOneByID(id int64) (*User, error)
}

type service struct {
	ur UserRepository
}

func NewUserService(ur UserRepository) UserService {
	return &service{ur: ur}
}

func (s *service) Create(user *User) (*User, error) {
	return s.ur.Create(user)
}
func (s *service) GetAll(role Role) ([]User, error) {
	return s.ur.GetAll(role)
}
func (s *service) GetOneByID(id int64) (*User, error) {
	return s.ur.GetOneByID(id)
}
