package models

import "time"

type User struct {
	ID        int       `json:"id" xml:"id"`
	FisrtName string    `json:"firstName" xml:"firstName"`
	LastName  string    `json:"lastName" xml:"lastName"`
	Username  string    `json:"username" xml:"username"`
	Email     string    `json:"email" xml:"email"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt"`
}
