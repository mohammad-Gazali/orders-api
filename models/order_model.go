package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uint64     `json:"id"`
	CustomerID  uuid.UUID  `json:"customer_id"`
	LineItems   []LineItem `json:"line_items"`
	CreatedAt   *time.Time `json:"created_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type LineItem struct {
	ID       uuid.UUID `json:"id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
