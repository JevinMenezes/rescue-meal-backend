package login

type Service interface {
	Login(User) string
}

type service struct{}

func (s service) Login(user User) (string, error) {
	return "success", nil
}
