package postgres

import (
	"github.com/malekelthomas/ShelterCM-API/pkg/user"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore() *UserStore {
	return &UserStore{
		db: gdb,
	}
}

func (us *UserStore) CreateUser(user *user.User) (*user.User, error) {
	result := us.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
