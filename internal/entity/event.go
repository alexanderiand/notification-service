package entity

import "time"

// Event entity
type Event struct {
	ID int `json:"id,omitempty"`
	OrderType string `json:"orderType"`
	SessionID string `json:"sessionId"`
	Card string `json:"card"`
	EventDate time.Time `json:"eventDate"`
	WebSiteURL string `json:"websiteUrl"`
	CreatedAt time.Time `json:"-"`
}

// TODO: GenFakeEvent, implement method for generating event with fake data