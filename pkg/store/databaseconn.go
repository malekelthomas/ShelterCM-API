package store

import (
	"database/sql"

	"gorm.io/gorm"
)

type DatabaseConn struct {
	DB *sql.DB
}

type GormDBConn struct {
	DB *gorm.DB
}
