package demo3

type GoogleAuthService struct {
	// endpoint ..etc
}

func (g GoogleAuthService) Verify(username, password string) (bool, error) {
	// implement ....
	return true, nil
}

type LoginService struct {
	AuthService GoogleAuthService
}

func (l LoginService) Log(username, password string) bool {
	ok, err := l.AuthService.Verify(username, password)
	if err != nil || ok == false {
		return false
	}
	return true
}

type IAuthService interface {
	Verify(username, password string) (bool, error)
}

type MockAuthService struct {
}

func (MockAuthService) Verify(username, password string) (bool, error) {
	return true, nil
}
