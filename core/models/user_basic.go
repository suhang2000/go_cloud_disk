package models

import "time"

type UserBasic struct {
	Id        int
	Identity  string
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
