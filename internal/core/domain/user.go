package domain

import "time"

type User struct {
	Id        string
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
}
