package signature

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Signature struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	SignatureType SignatureType  `json:"signature_type"`
	Source        string         `json:"source"`
}
type SignatureType string

const (
	SignatureTypeWritten SignatureType = "written"
	SignatureTypeQr      SignatureType = "qr"
)
