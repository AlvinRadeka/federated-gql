// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	UserType int    `json:"userType"`
}

func (User) IsEntity() {}
