package problem

import (
	"fmt"
	"os"
	"testing"

	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestNewProblem(t *testing.T) {
	problemId := "abc100_a"

	t.Run("正しいproblemIdを渡すとProblemが生成される", func(t *testing.T) {
		problem, err := NewProblem(problemId)

		if problem == nil {
			t.Errorf("problem should not be nil")
		}

		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}

		problem.DeleteProblemDir()
	})

	t.Run("不正なproblemIdを渡すとエラーが返る", func(t *testing.T) {
		problemId := "abc354"
		problem, err := NewProblem(problemId)

		if problem != nil {
			t.Errorf("problem should be nil")
		}

		if err == nil {
			t.Errorf("err should not be nil")
		}
	})

	t.Run("問題のディレクトリが作成される", func(t *testing.T) {
		problem, _ := NewProblem(problemId)

		// ディレクトリが作成されているか確認
		_, err := os.Stat(problem.ProblemDirPath())
		if err != nil {
			t.Errorf("problem directory should be created")
		}

		problem.DeleteProblemDir()
	})
}

func TestGetter(t *testing.T) {
	problemId := "abc100_a"
	contextName := "abc100"
	rank := "a"

	problemDirPath := fmt.Sprintf("%s/contests/abc100/a", setting.RootDir)

	problemUrl := "https://atcoder.jp/contests/abc100/tasks/abc100_a"
	submit_url := "https://atcoder.jp/contests/abc100/submissions/me"

	problem, _ := NewProblem(problemId)

	t.Run("ProblemId()でproblemIdが取得できる", func(t *testing.T) {
		if problem.ProblemId() != problemId {
			t.Errorf("problemId should be %s, but got %s", problemId, problem.ProblemId())
		}
	})

	t.Run("ContestName()でコンテスト名が取得できる", func(t *testing.T) {
		if problem.ContestName() != contextName {
			t.Errorf("contestName should be abc100, but got %s", problem.ContestName())
		}
	})

	t.Run("Rank()で問題のランクが取得できる", func(t *testing.T) {
		if problem.Rank() != rank {
			t.Errorf("rank should be a, but got %s", problem.Rank())
		}
	})

	t.Run("ProblemDir()で問題のディレクトリが取得できる", func(t *testing.T) {
		if problem.ProblemDirPath() != problemDirPath {
			t.Errorf("problemDirPath should be %s, but got %s", problemDirPath, problem.ProblemDirPath())
		}
	})

	t.Run("ProblemUrl()で問題閲覧ページのURLが取得できる", func(t *testing.T) {
		if problem.ProblemUrl() != problemUrl {
			t.Errorf("problemUrl should be %s, but got %s", problemUrl, problem.ProblemUrl())
		}
	})

	t.Run("SubmissionUrl()で問題提出ページのURLが取得できる", func(t *testing.T) {
		if problem.SubmissionUrl() != submit_url {
			t.Errorf("SubmissionUrl should be %s, but got %s", submit_url, problem.SubmissionUrl())
		}
	})
}

func TestCreateProblemDir(t *testing.T) {
	problemId := "abc100_a"
	problem, _ := NewProblem(problemId)

	t.Run("問題のディレクトリを作成できる", func(t *testing.T) {
		err := problem.CreateProblemDir()
		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}

		// ディレクトリが作成されているか確認
		_, err = os.Stat(problem.ProblemDirPath())
		if err != nil {
			t.Errorf("problem directory should be created")
		}
	})

	problem.DeleteProblemDir()
}

func TestDeleteProblemDir(t *testing.T) {
	problemId := "abc100_a"
	problem, _ := NewProblem(problemId)

	t.Run("問題のディレクトリを削除できる", func(t *testing.T) {
		problem.CreateProblemDir()

		err := problem.DeleteProblemDir()
		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}

		// ディレクトリが削除されているか確認
		_, err = os.Stat(problem.ProblemDirPath())
		if err == nil {
			t.Errorf("problem directory should be deleted")
		}
	})

	problem.DeleteProblemDir()
}

func TestValidateProblemId(t *testing.T) {
	okList := []string{
		"abc100_a",
		"abc1_a",
		"arc100_f",
		"agc100_f",
	}

	t.Run("正しいproblemIdを渡すとnilが変える", func(t *testing.T) {
		for _, problemId := range okList {
			err := validateProblemId(problemId)
			if err != nil {
				t.Errorf("problemId %s should be valid", problemId)
			}
		}
	})

	ngList := []string{
		"abc100_",
		"abc100_1",
		"abc100_abc",
		"arc100",
	}

	t.Run("不正なproblemIdを渡すとエラーが返る", func(t *testing.T) {
		for _, problemId := range ngList {
			err := validateProblemId(problemId)
			if err == nil {
				t.Errorf("problemId %s should be invalid", problemId)
			}
		}
	})
}
