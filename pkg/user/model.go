package user

import (
	"gorm.io/gorm"
)

type Role string

const (
	RoleStaff      Role = "staff"
	RoleSupervisor Role = "supervisor"
	RoleClient     Role = "client"
	RoleGuard      Role = "guard"
	RoleManager    Role = "manager"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      Role   `json:"role"`
}
