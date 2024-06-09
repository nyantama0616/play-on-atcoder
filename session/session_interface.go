package session

type ISession interface {
	Login() error
	Logout() error
	IsLoggedIn() bool
	SessionId() string
}
