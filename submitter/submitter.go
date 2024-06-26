package submitter

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gocolly/colly"
	"github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/session"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

type Submitter struct {
	problem    problem.IProblem
	session    session.ISession
	collector  *colly.Collector
	cookies    []*http.Cookie //TODO: cookieの管理場所を考える
	formData   map[string]string
	retryCount int
}

// SubmitterがISubmitterを実装していることを確認
var _ ISubmitter = (*Submitter)(nil)

// 新しいSubmitterを作成する
func NewSubmitter(problem problem.IProblem, session session.ISession) *Submitter {
	collector := colly.NewCollector()
	cookies := []*http.Cookie{
		{
			Name:  "REVEL_SESSION",
			Value: session.SessionId(),
		},
	}
	formData := make(map[string]string)

	return &Submitter{
		problem:    problem,
		session:    session,
		collector:  collector,
		cookies:    cookies,
		formData:   formData,
		retryCount: setting.APIMaxRetry,
	}
}

/*
ソースコードを提出する

	file: ソースコードのファイル

	提出に失敗にすると、エラーを返す
*/
func (s *Submitter) Submit(language string, file *os.File) error {
	sourceCode, _ := os.ReadFile(file.Name())

	success := false

	s.collector.OnHTML("form", func(e *colly.HTMLElement) {
		action := fmt.Sprintf("/contests/%s/submit", s.problem.ContestName())
		if !success && e.Attr("action") != action {
			return
		}

		actionUrl := fmt.Sprintf("%s%s", s.problem.RootUrl(), action)

		// // Fill in the form fields.

		// 言語を選択
		languageId := ""
		e.ForEach("select[name='data.LanguageId'] option", func(_ int, e *colly.HTMLElement) {
			if e.Text == language {
				fmt.Println("Language found!")
				languageId = e.Attr("value")
			}
		})

		csrf_token := e.ChildAttr("input[name='csrf_token']", "value")
		s.formData["data.TaskScreenName"] = s.problem.ProblemId()
		s.formData["data.LanguageId"] = languageId
		s.formData["sourceCode"] = string(sourceCode)
		s.formData["csrf_token"] = csrf_token

		s.collector.SetCookies(actionUrl, s.cookies)

		s.collector.OnHTMLDetach("form")

		s.collector.Post(actionUrl, s.formData)
	})

	s.collector.OnResponse(func(r *colly.Response) {
		if s.successSubmit(r) {
			success = true
		} else if r.Request.Method == "POST" {
			//リトライ
			s.retryCount--
			if s.retryCount > 0 {
				s.collector.Post(r.Request.URL.String(), s.formData)
			}
		}
	})

	url := s.problem.ProblemUrl()

	s.collector.SetCookies(url, s.cookies)

	err := s.collector.Visit(url)
	if err != nil {
		return fmt.Errorf("failed to visit: %v", err)
	}
	fmt.Printf("Visited %s\n", url)

	if !success {
		return fmt.Errorf("failed to submit")
	}

	return nil
}

func (s *Submitter) successSubmit(r *colly.Response) bool {
	return r.StatusCode == 200 && r.Request.URL.String() == s.problem.SubmissionUrl()
}
