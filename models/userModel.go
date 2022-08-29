package models

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

var (
	users = map[string]*User{"b7392d0b-8a6f-4436-869a-037d054ea7d5": {
		ID:        "b7392d0b-8a6f-4436-869a-037d054ea7d5",
		FirstName: "Segun",
		LastName:  "Bayo",
		Email:     "segs@gmail.com",
	}}
)

func (user *User) CreateUser() (User, error) {
	user.ID = uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user.Password = string(hashedPassword)
	users[user.ID] = user
	return *user, nil
}

func (user *User) GetUsers() (map[string]*User, error) {
	return users, nil
}

func (user *User) GetUser() (User, error) {
	mUser, isAvailable := users[user.ID]
	// err := bcrypt.CompareHashAndPassword([]byte(mUser.Password), []byte(user.Password))
	// if err != nil {
	// 	return User{}, err
	// }

	if !isAvailable {
		return User{}, errors.New("User not registered")
	}
	return *mUser, nil
}

func updateUser() {}

func deleteUser() {}
