package services

type MailService interface {
}

type MailServices struct {
}

func NewMailService() MailService {
	return &MailServices{}
}
