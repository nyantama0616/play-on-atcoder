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

func NewSession() *Session {
	c := colly.NewCollector()
	sessionIdPath := fmt.Sprintf("%s/secrets/session_id.txt", setting.RootDir)

	return &Session{
		collector:     c,
		sessionIdPath: sessionIdPath,
	}
}

func (s *Session) Login(username, password string) error {
	url := "https://atcoder.jp/login"

	s.collector.OnHTML("form.form-horizontal", func(e *colly.HTMLElement) {
		actionURL := e.Attr("action")

		// Fill in the form fields.
		formData := make(map[string]string)
		formData["username"] = username
		formData["password"] = password

		// Submit the form
		s.collector.Post(actionURL, formData)
	})

	s.collector.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200 {
			// クッキーを保存
			cookies := s.collector.Cookies(url)
			s.saveSessionId(cookies)
		}
	})

	// ログインページにアクセス
	err := s.collector.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return err
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

func (s *Session) IsLoggedIn() bool {
	return s.SessionId() != ""
}

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
