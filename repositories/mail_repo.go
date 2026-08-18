package repositories

import (
	"time"

	"github.com/ADMex1/GoProject/config"
	"github.com/ADMex1/GoProject/models"
)

type MailRepository interface {
	CreateMailThread(mailThread *models.MailThread) error
	CreateMailReceiver(mailReceiver *models.MailReceiver) error
	CreateMailAttachment(mailAttachment *models.MailAttachment) error
	FindThreadByPublicID(publicID *string) (*models.MailThread, error)
	FetchReceiverByThreadPublicID(mailThreadPublicID string) ([]models.User, error)
	AddReceiver(mailThreadID uint, userIDs []uint) error
	DeleteMailThread(id uint) error
	CreateMail(mail *models.Mail) error
}

type MailRepositorys struct {
}

func NewMailRepository() MailRepository {
	return &MailRepositorys{}
}

func (r *MailRepositorys) CreateMailThread(mailThread *models.MailThread) error {
	return config.DB.Create(mailThread).Error
}
func (r *MailRepositorys) CreateMailReceiver(mailReceiver *models.MailReceiver) error {
	return config.DB.Create(mailReceiver).Error
}

func (r *MailRepositorys) CreateMailAttachment(mailAttachment *models.MailAttachment) error {
	return config.DB.Create(mailAttachment).Error
}

func (r *MailRepositorys) FindThreadByPublicID(publicID *string) (*models.MailThread, error) {
	var mailThread models.MailThread
	err := config.DB.Where("public_id=?", publicID).First(&mailThread).Error
	return &mailThread, err
}

func (r *MailRepositorys) FetchReceiverByThreadPublicID(mailThreadPublicID string) ([]models.User, error) {
	var users []models.User
	err := config.DB.Joins("JOIN mail_receivers on mail_receivers.user_internal_id = users.internal_id").Joins("JOIN mail_threads on mail_threads.internal_id = mail_receivers.mail_thread_internal_id").Where("mail_threads.public_id = ?", mailThreadPublicID).Find(&users).Error
	return users, err
}

func (r *MailRepositorys) AddReceiver(mailThreadID uint, userIDs []uint) error {
	if len(userIDs) == 0 {
		return nil
	}
	now := time.Now()
	var receiver []models.MailReceiver
	for _, userIDs := range userIDs {
		receiver = append(receiver, models.MailReceiver{
			MailThreadID: int64(mailThreadID),
			UserID:       int64(userIDs),
			SentAt:       now,
		})
	}
	return config.DB.Create(&receiver).Error
}

func (r *MailRepositorys) DeleteMailThread(id uint) error {
	return config.DB.Delete(&models.MailThread{}, id).Error
}

func (r *MailRepositorys) CreateMail(mail *models.Mail) error {
	return config.DB.Create(mail).Error
}

// func (r *MailRepositorys) FindMailThreadByID(id uint) (*models.MailThread, error) {
// 	var mailThread models.MailThread
// 	err := config.DB.Preload("")
// }
