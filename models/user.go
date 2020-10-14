package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("New user must not include User Id")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUserById(userId int) (User, error) {
	for _, user := range users {
		if user.ID == userId {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", userId)
}

func UpdateUser(user User) (User, error) {
	for _, u := range users {
		if user.ID == u.ID {
			u.FirstName = user.FirstName
			u.LastName = user.LastName
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", user.ID)
}

func RemoveUserById(userId int) error {
	for i, user := range users {
		if user.ID == userId {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", userId)
}
