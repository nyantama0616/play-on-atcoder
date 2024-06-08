package cpp

import (
	"fmt"
	"os"
	"testing"

	. "github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestCompile(t *testing.T) {
	problem, _ := NewProblem("abc354_c")
	defer problem.RemoveProblemDir()

	executorCpp := NewExecutorCpp(
		problem,
		SourceCodePath{
			MainPath:       fmt.Sprintf("%s/executor/cpp/assets/main.cpp", setting.RootDir),
			IncludeDirPath: fmt.Sprintf("%s/executor/cpp/assets/include", setting.RootDir),
		},
	)

	err := executorCpp.Compile()
	if err != nil {
		t.Errorf("Compile() failed: %v", err)
	}

	t.Run("dest.cppがdest_expected.cppと等しい", func(t *testing.T) {
		destCppExpectedPath := fmt.Sprintf("%s/executor/cpp/assets/dest_expected.cpp", setting.RootDir)
		//２つのファイルを比較
		destCpp, _ := os.ReadFile(executorCpp.destCppPath)
		destCppExpected, _ := os.ReadFile(destCppExpectedPath)

		if string(destCpp) != string(destCppExpected) {
			t.Errorf("arrange() failed: dest.cpp is not same as expected")
		}
	})
}

func TestExecute(t *testing.T) {
	problem, _ := NewProblem("abc354_c")
	defer problem.RemoveProblemDir()

	executorCpp := NewExecutorCpp(
		problem,
		SourceCodePath{
			MainPath:       fmt.Sprintf("%s/executor/cpp/assets/main.cpp", setting.RootDir),
			IncludeDirPath: fmt.Sprintf("%s/executor/cpp/assets/include", setting.RootDir),
		},
	)

	executorCpp.Compile()

	inputFilePath := fmt.Sprintf("%s/executor/cpp/assets/input.txt", setting.RootDir)
	outputFilePath := fmt.Sprintf("%s/executor/cpp/assets/output.txt", setting.RootDir)

	inputFile, _ := os.Open(inputFilePath)
	outputFile, _ := os.Create(outputFilePath)
	defer inputFile.Close()
	defer outputFile.Close()
	defer os.Remove(outputFilePath)

	err := executorCpp.Execute(inputFile, outputFile, os.Stderr)
	if err != nil {
		t.Errorf("Execute() failed: %v", err)
	}

	t.Run("出力結果が正しい", func(t *testing.T) {
		output, _ := os.ReadFile(outputFilePath)
		expectedOutput := "Hello, panda!\n"

		if string(output) != expectedOutput {
			t.Errorf("Execute() failed: output is not same as expected")
		}
	})
}
