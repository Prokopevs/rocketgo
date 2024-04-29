package core

import (
	"context"
	"time"
)

type User struct {
	ID        int     
	FirstName string     
	UserName  string     
}

func (s *ServiceImpl) AddUser(ctx context.Context, user *User) error {
	u := user.toDB()
	now := time.Now()
	u.CreatedAt = &now

	return s.db.AddUser(ctx, u)
}

func (s *ServiceImpl) GetUser(ctx context.Context, id int) (*User, bool, error) {
	exists, err := s.db.IsUserWithIdExists(ctx, id)
	if err != nil {
		return nil, false, err
	}

	if !exists {
		return nil, true, ErrNoSuchUser
	}

	user, err := s.db.GetUser(ctx, id)
	if err != nil {
		return nil, false, err
	}

	return convertDBUserToService(user), true, nil
}

func (s *ServiceImpl) IsUserWithIdExists(ctx context.Context, id int) (bool, error) {
	return s.db.IsUserWithIdExists(ctx, id)
}