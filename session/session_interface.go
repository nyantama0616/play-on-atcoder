package session

type ISession interface {
	Login(string, string) error
	Logout() error
	IsLoggedIn() bool
	SessionId() string
}
