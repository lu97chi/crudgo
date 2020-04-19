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

func AddUsers(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, usr := range users {
		if usr.ID == id {
			return *usr, nil
		}
	}
	return User{}, fmt.Errorf("User with id '%v' not found", id)
}

func UpdateUser(usr User) (User, error) {
	for i, u := range users {
		if u.ID == usr.ID {
			users[i] = &usr
			return usr, nil
		}
	}
	return User{}, fmt.Errorf("User with id '%v' not found", usr.ID)
}

func RemoveUser(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id '%v' not found", id)
}
