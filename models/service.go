package models

import (
	"time"
	uuid "github.com/google/uuid"
)

type Service struct {
	ID          uuid.UUID `db:"srv_id" json:"token"`
	Name        string    `db:"srv_name" json:"name"`
	Description string    `db:"srv_description" json:"description"`
	Signature   string    `db:"srv_signature" json:"signature"`
	PrivateKey  string    `db:"srv_private_key" json:"private_key"`
	Status      int64     `db:"srv_status" json:"status"`
	CreatedAt   time.Time `db:"srv_created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"srv_updated_at" json:"updated_at"`
}
