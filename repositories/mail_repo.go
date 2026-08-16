package repositories

type MailRepository interface {
}

type MailRepositorys struct {
}

func NewMailRepository() MailRepository {
	return &MailRepositorys{}
}
