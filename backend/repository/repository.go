package repository

import (
	"fmt"
	"github.com/auth/backend/model"
)

type SqlRepository struct {
	users map[string]*model.User
	//session map[string]context.Se
}

func NewSqlRepository() (*SqlRepository, error) {
	repository := &SqlRepository{
		users: make(map[string]*model.User),
	}
	return repository, nil
}

func (r *SqlRepository) AddUser(user model.User)  {
	r.users[user.UserID] = &model.User{
		RawData:           user.RawData,
		Provider:          user.Provider,
		Email:             user.Email,
		Name:              user.Name,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		NickName:          user.NickName,
		Description:       user.Description,
		UserID:            user.UserID,
		AvatarURL:         user.AvatarURL,
		Location:          user.Location,
		AccessToken:       user.AccessToken,
		AccessTokenSecret: user.AccessTokenSecret,
		RefreshToken:      user.RefreshToken,
		ExpiresAt:         user.ExpiresAt,
		IDToken:           user.IDToken,
	}
}

func (r *SqlRepository) ExistsUser(userId string) bool {
	if r.users[userId] != nil {
		return true
	}
	return false
}

func (r *SqlRepository) GetUser(userId string) (*model.User, error) {
	if r.users[userId] != nil {
		return r.users[userId], nil
	}
	return nil, fmt.Errorf("invalid user on repository")
}

func (r *SqlRepository) GetUserView(userId string) (*model.UserView, error) {
	userFull := r.users[userId]
	if userFull != nil {
		return &model.UserView{
			ID:    userFull.UserID,
			Name:  userFull.Name,
			Image: userFull.AvatarURL,
			Email: userFull.Email,
		}, nil
	}
	return nil, fmt.Errorf("invalid user view on repository")
}
