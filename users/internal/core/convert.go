package core

import "github.com/Prokopevs/rocketgo/users/internal/pg"

func (n *User) toDB() *pg.User {
	return &pg.User{
		ID:        n.ID,
		FirstName: n.FirstName,
		UserName:  n.UserName,
	}
}

func convertDBUserToService(user *pg.User) *User {
	return &User{
		ID:        user.ID,
		FirstName: user.FirstName,
		UserName:  user.UserName,
	}
}