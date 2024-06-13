package mock

import (
	"fmt"
	"os"

	"github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

type MockProblem struct {
}

var _ problem.IProblem = (*MockProblem)(nil)

func NewMockProblem() *MockProblem {
	return &MockProblem{}
}

func (p *MockProblem) ProblemId() string {
	return "abc354_a"
}

func (p *MockProblem) ContestName() string {
	return "abc354"
}

func (p *MockProblem) Rank() string {
	return "a"
}

func (p *MockProblem) ProblemDirPath() string {
	return fmt.Sprintf("%s/contests/%s/%s", setting.RootDir, p.ContestName(), p.Rank())
}

func (p *MockProblem) CreateProblemDir() error {
	err := os.MkdirAll(p.ProblemDirPath(), 0755) //TODO: Permissionはこれでいいのか？
	if err != nil {
		return err
	}

	return nil
}

func (p *MockProblem) RemoveProblemDir() error {
	// problemディレクトリが存在する場合、problemディレクトリを削除する
	if _, err := os.Stat(p.ProblemDirPath()); err == nil {
		err := os.RemoveAll(p.ProblemDirPath())
		if err != nil {
			return err
		}
	}

	// contestディレクトリが存在し、かつ空の場合、contestディレクトリを削除する
	if _, err := os.Stat(p.ContestDirPath()); err == nil {
		files, err := os.ReadDir(p.ContestDirPath())
		if err != nil {
			return err
		}
		if len(files) == 0 {
			err := os.RemoveAll(p.ContestDirPath())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *MockProblem) ContestDirPath() string {
	return fmt.Sprintf("%s/contests/%s", setting.RootDir, p.ContestName())
}

func (p *MockProblem) ProblemUrl() string {
	return fmt.Sprintf("http://localhost:%d/problem", setting.MockServerPort)
}

func (p *MockProblem) SubmissionUrl() string {
	return fmt.Sprintf("http://localhost:%d/contests/abc354/submissions/me", setting.MockServerPort)
}

func (p *MockProblem) RootUrl() string {
	return fmt.Sprintf("http://localhost:%d", setting.MockServerPort)
}
