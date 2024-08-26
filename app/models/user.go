package models

import "database/sql"

type User struct {
	Id                   int
	Name                 string
	Email                string
	Password             string
	RememberToken        *string
	EmailVerifiedAt      sql.NullTime
	CreatedAt, UpdatedAt sql.NullTime
}
