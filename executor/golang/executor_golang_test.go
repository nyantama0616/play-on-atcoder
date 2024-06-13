package golang

import (
	"fmt"
	"os"
	"testing"

	"github.com/nyantama0616/play-on-atcoder/mock"
	"github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestArrange(t *testing.T) {
	problem := mock.NewMockProblem()
	defer problem.RemoveProblemDir()

	executorGolang := initExecutorGolang(problem)

	err := executorGolang.Arrange()
	if err != nil {
		t.Errorf("Arrange() failed: %v", err)
	}

	t.Run("dest.goがdest_expected.goと等しい", func(t *testing.T) {
		arrangedFile, _ := executorGolang.ArrangedFile()
		defer arrangedFile.Close()

		destGoPath := arrangedFile.Name()
		destGoExpectedPath := fmt.Sprintf("%s/executor/golang/_assets/dest_expected.go", setting.RootDir)

		//２つのファイルを比較
		destGo, _ := os.ReadFile(destGoPath)
		destGoExpected, _ := os.ReadFile(destGoExpectedPath)

		if string(destGo) != string(destGoExpected) {
			t.Errorf("arrange() failed: dest.go is not same as expected")
		}
	})
}

func TestCompile(t *testing.T) {
	problem := mock.NewMockProblem()
	defer problem.RemoveProblemDir()
	executorGolang := initExecutorGolang(problem)

	executorGolang.Arrange()

	err := executorGolang.Compile()
	if err != nil {
		t.Errorf("Compile() failed: %v", err)
	}
}

func TestExecute(t *testing.T) {
	problem := mock.NewMockProblem()
	defer problem.RemoveProblemDir()
	executorGolang := initExecutorGolang(problem)

	executorGolang.Arrange()
	executorGolang.Compile()

	inputFilePath := fmt.Sprintf("%s/executor/golang/_assets/input.txt", setting.RootDir)
	outputFilePath := fmt.Sprintf("%s/executor/golang/_assets/output.txt", setting.RootDir)

	inputFile, _ := os.Open(inputFilePath)
	outputFile, _ := os.Create(outputFilePath)
	defer inputFile.Close()
	defer outputFile.Close()
	defer os.Remove(outputFilePath)

	err := executorGolang.Execute(inputFile, outputFile, os.Stderr)
	if err != nil {
		t.Errorf("Execute() failed: %v", err)
	}

	t.Run("出力結果が正しい", func(t *testing.T) {
		output, _ := os.ReadFile(outputFilePath)
		expectedOutput := "6\n"

		if string(output) != expectedOutput {
			t.Errorf(fmt.Sprintf("output is %s, but expected %s", string(output), expectedOutput))
		}
	})
}

func initExecutorGolang(problem problem.IProblem) *ExecutorGolang {
	executorGolang := NewExecutorGolang(
		problem,
		SourceCodePath{
			MainPath: fmt.Sprintf("%s/executor/golang/_assets/main.go", setting.RootDir),
		},
	)

	return executorGolang
}
