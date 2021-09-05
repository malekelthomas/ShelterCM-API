package postgres

import (
	"sync"

	"github.com/malekelthomas/ShelterCM-API/pkg/user"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore() *UserStore {
	var o sync.Once
	o.Do(func() {
		gdb.AutoMigrate(&user.User{})
	})
	return &UserStore{
		db: gdb,
	}
}

func (us *UserStore) Create(user *user.User) (*user.User, error) {
	result := us.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
