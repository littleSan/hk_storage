package subMail

var _ Service = (*service)(nil)

type Service interface {
	i()
	Send(phone, content string) bool
	SendGlobal(phone, content string) bool
}

type service struct {
}

func New() *service {
	return &service{}
}

func (i *service) i() {
}
