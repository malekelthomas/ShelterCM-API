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
		if err := gdb.AutoMigrate(&user.User{}); err != nil {
			panic(err)
		}
	})
	return &UserStore{
		db: gdb,
	}
}

func (us *UserStore) Create(user *user.User) (*user.User, error) {
	result := us.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (us *UserStore) GetAll(role user.Role) ([]user.User, error) {
	var users []user.User
	result := us.db.Where(`role = ?`, role).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (us *UserStore) GetOneByID(id int64) (*user.User, error) {
	var user user.User
	result := us.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
