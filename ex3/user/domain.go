package user

import "time"

type User struct {
	ID        string
	Name      string
	Address   string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}
