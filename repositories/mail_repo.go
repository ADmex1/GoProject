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
	AddReceiver(mailThreadID uint, userIDs []uint) error
	DeleteMailThread(id uint) error
	CreateMail(mail *models.Mail) error

	FetchReceiverByThreadPublicID(mailThreadPublicID string) ([]models.User, error)
	FindThreadByPublicID(publicID *string) (*models.MailThread, error)
	FetchMailThreadsForReceiver(userPublicID, filter, sort string, limit, offset int) ([]models.MailThread, int64, error)
	FetchMailThreadForSender(userPublicID, filter, sort string, limit, offset int) ([]models.MailThread, int64, error)
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

// Fetching Zone
func (r *MailRepositorys) FetchReceiverByThreadPublicID(mailThreadPublicID string) ([]models.User, error) {
	var users []models.User
	err := config.DB.Joins("JOIN mail_receivers on mail_receivers.user_internal_id = users.internal_id").Joins("JOIN mail_threads on mail_threads.internal_id = mail_receivers.mail_thread_internal_id").Where("mail_threads.public_id = ?", mailThreadPublicID).Find(&users).Error
	return users, err
}

func (r *MailRepositorys) FetchMailThreadsForReceiver(userPublicID, filter, sort string, limit, offset int) ([]models.MailThread, int64, error) {
	var mailThreads []models.MailThread
	var total int64
	query := config.DB.Model(models.MailThread{}).Where("receiver_public_id = ? OR internal_id IN ("+"SELECT mail_receivers.mail_threads.internal_id = mail_receivers.user_internal_id"+"WHERE users.public_id = ?)", userPublicID, userPublicID)
	if filter != "" {
		query = query.Where("title ILIKE ?", "%"+filter+"%")
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if sort != "" {
		query = query.Order(sort)
	} else {
		query = query.Order("created_at desc")
	}
	if err := query.Limit(limit).Offset(offset).Find(mailThreads).Error; err != nil {
		return nil, 0, err
	}
	return mailThreads, total, nil

}

func (r *MailRepositorys) FetchMailThreadForSender(userPublicID, filter, sort string, limit, offset int) ([]models.MailThread, int64, error) {
	var mailThreads []models.MailThread
	var total int64
	query := config.DB.Model(&models.MailThread{}).Where("sender_public_id = ?", userPublicID)
	if filter != "" {
		query = query.Where("title ILIKE ?", "%"+filter+"%")
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if sort != "" {
		query = query.Order(sort)
	} else {
		query = query.Order("created_at desc")
	}
	if err := query.Limit(limit).Offset(offset).Find(mailThreads).Error; err != nil {
		return nil, 0, err
	}
	return mailThreads, total, nil

}
