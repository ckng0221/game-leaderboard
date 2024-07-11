package models

import "gorm.io/gorm"

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

func (r Role) String() string {
	return string(r)
}

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique" json:"username"`
	// Email    string `gorm:"unique"`
	// Password string `gorm:"type:varchar(255)" json:"-"`
	Role Role `gorm:"type:enum('admin', 'player'); default:player" json:"role"`
}

var Roles = [...]Role{Admin, Member}
