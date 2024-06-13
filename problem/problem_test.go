/* LookAtMe-2:
interfaceを定義したら、そのパッケージのテストを書く

テストを書くことで、そのパッケージが正しく動作するか確認できる
*/

package problem

import (
	"fmt"
	"os"
	"testing"

	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestNewProblem(t *testing.T) {

	t.Run("正しいproblemIdを渡すとProblemが生成される", func(t *testing.T) {
		problemId := "abc100_a"
		problem, err := NewProblem(problemId)
		defer problem.RemoveProblemDir()

		if problem == nil {
			t.Errorf("problem should not be nil")
		}

		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}
	})

	t.Run("不正なproblemIdを渡すとエラーが返る", func(t *testing.T) {
		problemId := "abc100"
		problem, err := NewProblem(problemId)

		if problem != nil {
			t.Errorf("problem should be nil")
		}

		if err == nil {
			t.Errorf("err should not be nil")
		}
	})

	t.Run("問題のディレクトリが作成される", func(t *testing.T) {
		problemId := "abc100_a"
		problem, _ := NewProblem(problemId)
		defer problem.RemoveProblemDir()

		// ディレクトリが作成されているか確認している
		_, err := os.Stat(problem.ProblemDirPath())
		if err != nil {
			t.Errorf("problem directory should be created")
		}
	})
}

func TestGetter(t *testing.T) {
	problemId := "abc100_a"
	problem, _ := NewProblem(problemId)
	defer problem.RemoveProblemDir()

	t.Run("ProblemId()でproblemIdが取得できる", func(t *testing.T) {
		expected := "abc100_a"
		if problem.ProblemId() != expected {
			t.Errorf("problemId should be %s, but got %s", expected, problem.ProblemId())
		}
	})

	t.Run("ContestName()でコンテスト名が取得できる", func(t *testing.T) {
		expected := "abc100"
		if problem.ContestName() != expected {
			t.Errorf("contestName should be %s, but got %s", expected, problem.ContestName())
		}
	})

	t.Run("Rank()で問題のランクが取得できる", func(t *testing.T) {
		expected := "a"
		if problem.Rank() != expected {
			t.Errorf("rank should be %s, but got %s", expected, problem.Rank())
		}
	})

	t.Run("ProblemDirPath()で問題のディレクトリが取得できる", func(t *testing.T) {
		expected := fmt.Sprintf("%s/contests/abc100/a", setting.RootDir)
		if problem.ProblemDirPath() != expected {
			t.Errorf("problemDirPath should be %s, but got %s", expected, problem.ProblemDirPath())
		}
	})

	t.Run("ProblemUrl()で問題閲覧ページのURLが取得できる", func(t *testing.T) {
		expected := "https://atcoder.jp/contests/abc100/tasks/abc100_a"
		if problem.ProblemUrl() != expected {
			t.Errorf("problemUrl should be %s, but got %s", expected, problem.ProblemUrl())
		}
	})

	t.Run("SubmissionUrl()で問題提出ページのURLが取得できる", func(t *testing.T) {
		expected := "https://atcoder.jp/contests/abc100/submissions/me"
		if problem.SubmissionUrl() != expected {
			t.Errorf("SubmissionUrl should be %s, but got %s", expected, problem.SubmissionUrl())
		}
	})

	t.Run("RootUrl()で問題情報を取得するサーバのルートURLが取得できる", func(t *testing.T) {
		expected := "https://atcoder.jp"
		if problem.RootUrl() != expected {
			t.Errorf("RootUrl should be %s, but got %s", expected, problem.RootUrl())
		}
	})
}

func TestCreateProblemDir(t *testing.T) {
	problemId := "abc100_a"
	problem, _ := NewProblem(problemId)
	defer problem.RemoveProblemDir()

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
}

func TestRemoveProblemDir(t *testing.T) {
	problemId := "abc100_a"
	problem, _ := NewProblem(problemId)
	problem.RemoveProblemDir()

	t.Run("問題のディレクトリを削除できる", func(t *testing.T) {
		problem.CreateProblemDir()
		defer problem.RemoveProblemDir()

		err := problem.RemoveProblemDir()
		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}

		// ディレクトリが削除されているか確認
		_, err = os.Stat(problem.ProblemDirPath())
		if err == nil {
			t.Errorf("problem directory should be deleted")
		}
	})

	t.Run("問題のディレクトリが存在しない場合、何もしない", func(t *testing.T) {
		// ディレクトリが存在するか確認
		_, err := os.Stat(problem.ProblemDirPath())
		if err == nil {
			t.Errorf("problem directory should not be created")
		}

		err = problem.RemoveProblemDir()
		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}
	})
}

func TestValidateProblemId(t *testing.T) {
	okList := []string{
		"abc100_a",
		"abc1_a",
		"arc100_f",
		"agc100_f",
	}

	t.Run("正しいproblemIdを渡すとnilが返る", func(t *testing.T) {
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
