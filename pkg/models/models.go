package models

import (
	"net/http"
)

type Event struct {
	ID        string `json:"id"`
	Status    string `json:"order_date"`
	ForceStop string `json:"force_stop"`
	CreatedAt string `json:created_at`
	UpdatedAt string `json:updated_at`
	JobType   string `json:job_type`
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
