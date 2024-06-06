package problem

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/nyantama0616/play-on-atcoder/setting"
)

type Problem struct {
	problemId      string
	contestName    string
	rank           string
	contestDirPath string
}

// ProblemがIProblemを実装していることをコンパイル時に確認
var _ IProblem = (*Problem)(nil)

// コンストラクタ
func NewProblem(problemId string) (*Problem, error) {

	err := validateProblemId(problemId)
	if err != nil {
		return nil, err
	}

	split := strings.Split(problemId, "_")

	contestName := split[0]
	rank := split[1]
	contestDirPath := fmt.Sprintf("%s/contests/%s", setting.RootDir, contestName)

	problem := &Problem{
		problemId:      problemId,
		contestName:    contestName,
		rank:           rank,
		contestDirPath: contestDirPath,
	}

	problem.CreateProblemDir()

	return problem, nil
}

func (p *Problem) ProblemId() string {
	return p.problemId
}

func (p *Problem) ContestName() string {
	return p.contestName
}

func (p *Problem) Rank() string {
	return p.rank
}

func (p *Problem) ContestDirPath() string {
	return fmt.Sprintf("%s/contests/%s", setting.RootDir, p.ContestName())
}

func (p *Problem) ProblemDirPath() string {
	return fmt.Sprintf("%s/%s", p.ContestDirPath(), p.Rank())
}

func (p *Problem) ProblemUrl() string {
	return fmt.Sprintf("https://atcoder.jp/contests/%s/tasks/%s", p.ContestName(), p.ProblemId())
}

func (p *Problem) SubmissionUrl() string {
	return fmt.Sprintf("https://atcoder.jp/contests/%s/submissions/me", p.ContestName())
}

func (p *Problem) CreateProblemDir() error {
	err := os.MkdirAll(p.ProblemDirPath(), 0755) //TODO: Permissionはこれでいいのか？
	if err != nil {
		return err
	}

	return nil
}

func (p *Problem) DeleteProblemDir() error {
	err := os.RemoveAll(p.ProblemDirPath())
	if err != nil {
		return err
	}

	// contestディレクトリが空の場合、contestディレクトリを削除する
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

	return nil
}

func validateProblemId(problemId string) error {

	// problemIdは"abc000_x"の形式である必要がある
	if !regexp.MustCompile(`^(abc|arc|agc)[0-9]{1,3}_[a-z]$`).MatchString(problemId) {
		return errors.New("problemId is invalid")
	}

	return nil
}
