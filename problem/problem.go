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

/*
新しいProblemを生成する

	problemIdが不正な場合はエラーを返す
*/
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

/*
問題IDを取得する

	例: "abc100_a"
*/
func (p *Problem) ProblemId() string {
	return p.problemId
}

/*
コンテスト名を取得する

	例: "abc100"
*/
func (p *Problem) ContestName() string {
	return p.contestName
}

/*
問題の難易度を取得する

	例: "a"
*/
func (p *Problem) Rank() string {
	return p.rank
}

/*
問題のディレクトリのパスを取得する

	例: "contests/abc100/a"
*/
func (p *Problem) ContestDirPath() string {
	return fmt.Sprintf("%s/contests/%s", setting.RootDir, p.ContestName())
}

/*
問題のディレクトリのパスを取得する

	例: "contests/abc100/a"
*/
func (p *Problem) ProblemDirPath() string {
	return fmt.Sprintf("%s/%s", p.ContestDirPath(), p.Rank())
}

/*
問題のURLを取得する

	例: "https://atcoder.jp/contests/abc100/tasks/abc100_a"
*/
func (p *Problem) ProblemUrl() string {
	return fmt.Sprintf("https://atcoder.jp/contests/%s/tasks/%s", p.ContestName(), p.ProblemId())
}

/*
提出ページのURLを取得する

	例: "https://atcoder.jp/contests/abc100/submissions/me"
*/
func (p *Problem) SubmissionUrl() string {
	return fmt.Sprintf("https://atcoder.jp/contests/%s/submissions/me", p.ContestName())
}

/*
問題ディレクトリを作成する

	問題ディレクトリが既に存在する場合は何もしない

	例: "contests/abc100/a"
*/
func (p *Problem) CreateProblemDir() error {
	err := os.MkdirAll(p.ProblemDirPath(), 0755) //TODO: Permissionはこれでいいのか？
	if err != nil {
		return err
	}

	return nil
}

/*
問題ディレクトリが存在する場合、問題ディレクトリを削除する。

	存在しない場合は、何もしない

	例: "contests/abc100/a"
*/
func (p *Problem) RemoveProblemDir() error {
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

/*
problemIdの形式を検証する

	problemIdは"abc000_x"の形式である必要がある
	不正な形式の場合はエラーを返す
*/
func validateProblemId(problemId string) error {

	if !regexp.MustCompile(`^(abc|arc|agc)[0-9]{1,3}_[a-z]$`).MatchString(problemId) {
		return errors.New("problemId is invalid")
	}

	return nil
}
