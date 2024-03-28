package domain

import "time"

type User struct {
	id        string
	name      string
	username  string
	password  string
	createdAt time.Time
}
