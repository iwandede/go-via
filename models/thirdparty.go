package models

import (
	"time"
)

type ThirdParty struct {
	ID          string    `db:"trp_id" json:"id"`
	Code        string    `db:"trp_code" json:"code"`
	Name        string    `db:"trp_name" json:"name"`
	Description string    `db:"trp_description" json:"description"`
	URL         string    `db:"trp_url" json:"url"`
	Command     string    `db:"trp_command" json:"command"`
	Status      int64     `db:"trp_status" json:"status"`
	CreatedAt   time.Time `db:"trp_created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"trp_updated_at" json:"updated_at"`
}

type ResponseThirdParty struct {
	Status      string `json:"status"`
	Messageid   string `json:"messageid"`
	Destination string `json:"destination"`
}