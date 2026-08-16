package models

import "time"

type MailReceiver struct {
	MailReceiverID int64     `json:"mail_receiver_id" gorm:"column:mail_receiver_internal_id;primaryKey;autoIncrement:false"`
	MailThreadID   int64     `json:"mail_thread_id" gorm:"column:mail_thread_internal_id;primaryKey;autoIncrement:false"`
	UserID         int64     `json:"user_id" gorm:"column:user_internal_id;primaryKey;autoIncrement:false"`
	SentAt         time.Time `json:"sent_at" db:"sent_at" gorm:"autoCreateTime"`
}
