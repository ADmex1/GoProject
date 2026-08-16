package repositories

import (
	"github.com/ADMex1/GoProject/config"
	"github.com/ADMex1/GoProject/models"
)

type MailRepository interface {
}

type MailRepositorys struct {
}

func NewMailRepository() MailRepository {
	return &MailRepositorys{}
}

func (r *MailRepositorys) CreateMailThread(MailThread *models.MailThread) error {
	return config.DB.Create(MailThread).Error
}

// func (r *MailRepositorys) FindMailThreadByID(id uint) (*models.MailThread,error){
// 	var MailThread models.MailThread
// 	err := config.DB.Preload("")
// }
