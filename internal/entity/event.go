package entity

import "time"

// Event entity
type Event struct {
	ID         int       `json:"-" sql:"id"`
	OrderType  string    `json:"orderType" sql:"order_type"`
	SessionID  string    `json:"sessionId" sql:"session_id"`
	Card       string    `json:"card" sql:"card"`
	EventDate  string    `json:"eventDate" sql:"event_date"`
	WebSiteURL string    `json:"websiteUrl" sql:"website_url"`
	CreatedAt  time.Time `json:"-" sql:"created_at"`
}

// TODO: GenFakeEvent, implement method for generating event with fake data
