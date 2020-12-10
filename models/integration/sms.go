package integration

import "time"

type SMSRequest struct {
	PhoneNumber string `json:"phone_number"`
	Message     string `json:"message"`
	DeviceID    int    `json:"device_id"`
	From        string `json:"from,omitempty"`
}

type SMSResponse struct {
	ID          int       `json:"id"`
	DeviceID    int       `json:"device_id"`
	PhoneNumber string    `json:"phone_number"`
	Message     string    `json:"message"`
	Status      string    `json:"status"`
	Log         []*Log    `json:"log"`
	CreatedAt   time.Time `json:"created_at"`
}

type Log struct {
	Status     string    `json:"status"`
	OccurredAt time.Time `json:"occurred_at"`
}
