package entity

import "time"

// Client entity
type Client struct {
	ID int
	Username string
	FirstName string
	LastName string
	DateOfBirth time.Time
	Email string
	Password string
	AccountID int
}

// TODO: GenFakeClient, for testing service