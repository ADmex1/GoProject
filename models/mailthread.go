package models

import (
	"time"

	"github.com/google/uuid"
)

type MailThread struct {
	InternalID     int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID       uuid.UUID `json:"public_id" db:"public_id" `
	SenderID       int64     `json:"sender_internal_id" gorm:"column:sender_internal_id"`
	SenderPublicID uuid.UUID `json:"sender_public_id" db:"sender_public_id"`
	Subject        string    `json:"subject" db:"subject" gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
