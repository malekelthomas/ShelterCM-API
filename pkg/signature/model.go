package signature

import "gorm.io/gorm"

type Signature struct {
	gorm.Model
	UserID int64 `json:"user_id"`
}
