package model

import (
	"time"
)

// User represents the schema of table "users".
type User struct {
	Acct       string    `json:"acct"`
	Pwd        string    `json:"pwd"`
	Fullname   string    `json:"fullname"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Users []User

var UserModel User

func (u *User) TableName() string {
	return "users"
}

// func (u *User) Create() {
// }
