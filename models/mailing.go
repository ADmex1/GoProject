package models

import (
	"time"

	"github.com/google/uuid"
)

// type Mail struct {
// 	InternalID   int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
// 	PublicID     uuid.UUID `json:"public_id" db:"public_id" `
// 	MailThreadID int64     `json:"mail_thread_id" db:"mail_thread_id"`
// 	UserID       int64     `json:"user_internal_id" db:"user_internal_id"`
// 	UserPubID    uuid.UUID `json:"user_id" db:"user_id"`
// 	Message      string    `json:"message" db:"message" gorm:"type:text;not null"`

//		CreatedAt time.Time `json:"created_at" db:"created_at"`
//	}
type Mail struct {
	InternalID   int64     `json:"internal_id" gorm:"column:internal_id;primaryKey;autoIncrement"`
	PublicID     uuid.UUID `json:"public_id" gorm:"column:public_id;type:uuid;not null"`
	MailThreadID int64     `json:"mail_thread_id" gorm:"column:mail_thread_internal_id;not null"`
	UserID       int64     `json:"user_id" gorm:"column:user_internal_id;not null"`
	UserPubID    uuid.UUID `json:"user_pub_id" gorm:"column:user_public_id;type:uuid;not null"`
	Message      string    `json:"message" gorm:"column:message;type:text;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
