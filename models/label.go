package models

type Label struct {
	InternalID int64  `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   string `json:"public_id" db:"public_id" `
	LabelName  string `json:"label_name" db:"label_name"`
	Color      string `json:"color" db:"color"`
}
