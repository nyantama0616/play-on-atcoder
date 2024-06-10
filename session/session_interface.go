package session

type ISession interface {
	// ログインする
	Login() error

	// ログアウトする
	Logout() error

	// ログインしているかどうかを返す
	IsLoggedIn() bool

	// セッションIDを返す
	SessionId() string
}
