package services

import (
	"errors"

	"github.com/ADMex1/GoProject/models"
	"github.com/ADMex1/GoProject/repositories"
	"github.com/google/uuid"
)

type MailService interface {
}

type MailServices struct {
	userRepo repositories.UserRepository
	mailRepo repositories.MailRepository
}

func NewMailService() MailService {
	return &MailServices{}
}

func (s *MailServices) CreateNewMailThread(mailThread models.MailThread) error {
	users, err := s.userRepo.FindByPublicID(string(mailThread.SenderPublicID.String()))
	if err != nil {
		return errors.New("User not exist")
	}
	mailThread.PublicID = uuid.New()
	mailThread.SenderID = users.InternalID
	return s.mailRepo.CreateMailThread(&mailThread)
}

func (s *MailServices) AddReceiver(mailThreadPublicID string, userPublicIDs []string) error {
	mailThread, err := s.mailRepo.FindThreadByPublicID(&mailThreadPublicID)
	if err != nil {
		return errors.New("Mail Thread Not Found!")
	}
	var userInternalIDs []uint
	for _, userPublicID := range userPublicIDs {
		user, err := s.userRepo.FindByPublicID(userPublicID)
		if err != nil {
			return errors.New("User not found!")
		}
		userInternalIDs = append(userInternalIDs, uint(user.InternalID))
	}
	existingReceiver, err := s.mailRepo.FetchReceiverByThreadPublicID(string(mailThread.PublicID.String()))
	if err != nil {
		return err
	}

	receiverMap := make(map[uint]bool)
	for _, receiver := range existingReceiver {
		receiverMap[uint(receiver.InternalID)] = true
	}

	var NewReceiverIDs []uint
	for _, userID := range userInternalIDs {
		if !receiverMap[userID] {
			NewReceiverIDs = append(NewReceiverIDs, userID)
		}
	}
	if len(NewReceiverIDs) == 0 {
		return nil
	}
	return s.mailRepo.AddReceiver(uint(mailThread.InternalID), NewReceiverIDs)
}

func (s *MailServices) CreateNewMail(mail models.Mail) error {
	users, err := s.userRepo.FindByPublicID((mail.UserPubID).String())
	if err != nil {
		return errors.New("user not exist")
	}
	mail.PublicID = uuid.New()
	mail.UserID = users.InternalID
	return s.mailRepo.CreateMail(&mail)
}

func (s *MailServices) DeleteMailThread(id uint) error {
	return s.mailRepo.DeleteMailThread(id)
}
