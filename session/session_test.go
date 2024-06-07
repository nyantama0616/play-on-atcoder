package session

import (
	"testing"
)

func TestLogin(t *testing.T) {
	s := NewSession()
	defer s.Logout()

	username := "test16test"
	password := "kj9JwNnq003KgqJ"

	err := s.Login(username, password)

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
