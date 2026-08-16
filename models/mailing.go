package models

import (
	"time"

	"github.com/google/uuid"
)

type Mail struct {
	InternalID   int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID     uuid.UUID `json:"public_id" db:"public_id" `
	MailThreadID int64     `json:"mail_thread_id" db:"mail_thread_id"`
	UserID       int64     `json:"user_internal_id" db:"user_internal_id"`
	UserPubID    uuid.UUID `json:"user_id" db:"user_id"`
	Message      string    `json:"message" db:"message" gorm:"type:text;not null"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
