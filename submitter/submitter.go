package submitter

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gocolly/colly"
	"github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/session"
)

type Submitter struct {
	problem   problem.IProblem
	session   session.ISession
	collector *colly.Collector
	cookies   []*http.Cookie //TODO: cookieの管理場所を考える
}

func NewSubmitter(problem problem.IProblem, session session.ISession) *Submitter {
	collector := colly.NewCollector()
	cookies := []*http.Cookie{
		{
			Name:  "REVEL_SESSION",
			Value: session.SessionId(),
		},
	}

	return &Submitter{
		problem:   problem,
		session:   session,
		collector: collector,
		cookies:   cookies,
	}
}

func (s *Submitter) Submit(file *os.File) error {
	url := s.problem.ProblemUrl()

	sourceCode, _ := os.ReadFile(file.Name())

	success := false

	s.collector.OnHTML("form", func(e *colly.HTMLElement) {
		action := fmt.Sprintf("/contests/%s/submit", s.problem.ContestName())
		if !success && e.Attr("action") != action {
			return
		}

		actionUrl := fmt.Sprintf("https://atcoder.jp%s", action)
		fmt.Println(actionUrl)

		// // Fill in the form fields.
		formData := make(map[string]string)
		csrf_token := e.ChildAttr("input[name='csrf_token']", "value")
		formData["data.TaskScreenName"] = s.problem.ProblemId()
		formData["data.LanguageId"] = "5001"
		formData["sourceCode"] = string(sourceCode)
		formData["csrf_token"] = csrf_token

		s.collector.SetCookies(actionUrl, s.cookies)

		s.collector.OnHTMLDetach("form")

		err := s.collector.Post(actionUrl, formData)
		if err != nil {
			fmt.Printf("Failed to post: %v\n", err)
		} else {
			success = true
		}
	})

	s.collector.SetCookies(url, s.cookies)

	err := s.collector.Visit(url)
	if err != nil {
		return fmt.Errorf("failed to visit: %v", err)
	}
	fmt.Printf("Vosited %s\n", url)

	if !success {
		return fmt.Errorf("failed to submit")
	}

	return nil
}
