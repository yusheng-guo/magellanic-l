package service

type UserServiceT struct {
}

var UserService = new(UserServiceT)

// Login 登录
func (s *UserServiceT) Login(email, password string) (string, error) {
	return "", nil
}

// Register 注册
func (s *UserServiceT) Register() (string, error) {
	return "", nil
}
