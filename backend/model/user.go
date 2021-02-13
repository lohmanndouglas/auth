package model

import (
	"time"
)

// the user that we should send to front next-auth session
type UserView struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Email string `json:"email"`
}

// the user that we should have in the database
type User struct {
	RawData           map[string]interface{}
	Provider          string
	Email             string
	Name              string
	FirstName         string
	LastName          string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}

// a github model of user
type GithubUser struct {
	ID string `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	Image string `json:"image"`
	Email string `json:"email"`
}

// a google model of user
type GoogleUser struct {
	ID string `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	Image string `json:"image"`
	Email string `json:"email"`
}