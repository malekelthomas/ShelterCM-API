package postgres

import (
	"sync"

	"github.com/malekelthomas/ShelterCM-API/pkg/signature"
	"gorm.io/gorm"
)

type SignatureStore struct {
	db *gorm.DB
}

func NewSignatureStore() *SignatureStore {
	var o sync.Once
	o.Do(func() {
		gdb.AutoMigrate(&signature.Signature{})
	})
	return &SignatureStore{
		db: gdb,
	}
}

func (ss *SignatureStore) Create(signature *signature.Signature) (*signature.Signature, error) {
	result := ss.db.Create(&signature)
	if result.Error != nil {
		return nil, result.Error
	}

	return signature, nil
}
