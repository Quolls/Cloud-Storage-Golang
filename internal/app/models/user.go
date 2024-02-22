package models

import "database/sql"

type User struct {
	Username       string         `json:"user_name"`
	Password       string         `json:"user_pwd"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	EmailValidated int8           `json:"email_validated"`
	PhoneValidated int8           `json:"phone_validated"`
	SignUpAt       string         `json:"signup_at"`
	LastActive     string         `json:"last_active"`
	Profile        sql.NullString `json:"profile"`
	Status         int            `json:"status"`
}
