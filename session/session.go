package session

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

type Session struct {
	collector     *colly.Collector
	sessionIdPath string
}

// SessionがISessionを実装していることを確認
var _ ISession = (*Session)(nil)

// 新しいSessionを作成する
func NewSession() *Session {
	c := colly.NewCollector()
	c.AllowURLRevisit = true // 同じURLに何度もアクセスすることを許可(複数回のログイン試行を可能にするため)

	sessionIdPath := fmt.Sprintf("%s/secrets/session_id.txt", setting.RootDir)

	return &Session{
		collector:     c,
		sessionIdPath: sessionIdPath,
	}
}

func (s *Session) Login() error {
	var username string
	var password string

	//ユーザ名をScan
	fmt.Printf("Enter your username: ")
	fmt.Scan(&username)

	//パスワードをScan
	fmt.Printf("Enter your password: ")
	fmt.Scan(&password)

	return s.login(username, password)
}

// TODO: ちゃんと実装する, session_idを削除しただけではログアウトにならない
func (s *Session) Logout() error {
	// ファイルを削除
	err := os.Remove(s.sessionIdPath)
	if err != nil {
		return err
	}

	return nil
}

// ログインしているかどうかを返す
func (s *Session) IsLoggedIn() bool {
	return s.SessionId() != ""
}

// セッションIDを返す
func (s *Session) SessionId() string {
	file, err := os.Open(s.sessionIdPath)
	if err != nil {
		return ""
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		return ""
	}

	return string(buf[:n])
}

// 環境変数を読み込んでログインする
func (s *Session) LoginWithEnv() error {
	username := os.Getenv("ATCODER_USERNAME")
	password := os.Getenv("ATCODER_PASSWORD")

	return s.login(username, password)
}

func (s *Session) login(username, password string) error {
	url := "https://atcoder.jp/login"

	success := false
	s.collector.OnHTML("form", func(e *colly.HTMLElement) {
		if e.Attr("action") == "" {
			// Fill in the form fields.
			csrfToken := e.ChildAttr("input[name='csrf_token']", "value")
			formData := make(map[string]string)
			formData["username"] = username
			formData["password"] = password
			formData["csrf_token"] = csrfToken

			s.collector.OnHTMLDetach("form")

			// Submit the form
			err := s.collector.Post("https://atcoder.jp/login", formData)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	s.collector.OnResponse(func(r *colly.Response) {
		fmt.Println("response")
		fmt.Println(r.Request.URL.String())
		fmt.Println(r.Request.Method)
		fmt.Println(r.StatusCode)

		if s.successLogin(r) {
			// クッキーを保存
			cookies := s.collector.Cookies(url)
			s.saveSessionId(cookies)
			success = true
		}
	})

	// ログインページにアクセス
	err := s.collector.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	if !success {
		return fmt.Errorf("failed to login")
	}

	return err
}

// TODO: session_idのみ保存しているが、cookieを全部保存するようにするかもしれない
func (s *Session) saveSessionId(cokkies []*http.Cookie) error {
	// ファイルを開く
	file, err := os.Create(s.sessionIdPath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, cookie := range cokkies {

		if cookie.Name == "REVEL_SESSION" {
			// session_idを保存
			_, err := file.WriteString(cookie.Value)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("invalid cookies")
}

// ログインが成功したかどうかを返す
func (s *Session) successLogin(r *colly.Response) bool {
	return r.StatusCode == 200 && r.Request.URL.String() == "https://atcoder.jp/home"
}
