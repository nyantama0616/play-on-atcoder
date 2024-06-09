package session

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestLogin(t *testing.T) {
	envFilePath := fmt.Sprintf("%s/.env", setting.RootDir)
	if err := godotenv.Load(envFilePath); err != nil {
		panic("No .env file found")
	}

	s := NewSession()

	username := os.Getenv("ATCODER_USERNAME")
	password := os.Getenv("ATCODER_PASSWORD")

	err := s.login(username, password)

	t.Run("ログインが成功する", func(t *testing.T) {
		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}
	})

	t.Run("IsLoggedIn()がtrueになる", func(t *testing.T) {
		if !s.IsLoggedIn() {
			t.Errorf("IsLoggedIn should be true")
		}
	})

	t.Run("セッションIDが取得できる", func(t *testing.T) {
		sessionId := s.SessionId()
		if sessionId == "" {
			t.Errorf("sessionId should not be empty")
		}
	})
}
