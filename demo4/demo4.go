package demo4

type GoogleAuthService struct {
	// endpoint ..etc
}

func (g GoogleAuthService) Verify(username, password string) (bool, error) {
	// implementation ....
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
