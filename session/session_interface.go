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

	// 環境変数を読み込んでログインする
	LoginWithEnv() error
}
