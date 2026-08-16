package models

import (
	"time"

	"github.com/google/uuid"
)

type MailAttachment struct {
	InternalID int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" gorm:"type:uuid;not null"`
	// CardID     int64     `json:"card_internal_id" db:"card_internal_id" gorm:"column:card_internal_id"`
	MailID    int64     `json:"mail_internal_id" db:"mail_internal_id" gorm:"column:mail_internal_id"`
	UserID    int64     `json:"user_internal_id" db:"user_internal_id" gorm:"column:user_internal_id"`
	File      string    `json:"file" db:"file"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	FileURL string `json:"file_url" gorm:"-"`
}

func (MailAttachment) TableName() string {
	return "mail_attachment"
}
