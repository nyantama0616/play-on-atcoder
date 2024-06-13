package golang

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/nyantama0616/play-on-atcoder/executor"
	"github.com/nyantama0616/play-on-atcoder/problem"
)

type ExecutorGolang struct {
	outputDirPath  string
	destGoPath     string
	sourceCodePath SourceCodePath
}

type SourceCodePath struct {
	MainPath string
}

// IExecutorを実装しているか確認
var _ executor.IExecutor = (*ExecutorGolang)(nil)

func NewExecutorGolang(problem problem.IProblem, sourceCodePath SourceCodePath) *ExecutorGolang {
	outputDirPath := fmt.Sprintf("%s/executor/golang", problem.ProblemDirPath())
	destGoPath := fmt.Sprintf("%s/dest.go", outputDirPath)

	os.MkdirAll(outputDirPath, 0777)

	return &ExecutorGolang{
		outputDirPath:  outputDirPath,
		destGoPath:     destGoPath,
		sourceCodePath: sourceCodePath,
	}
}

/*
提出可能な形にソースコードを整形する。
現状はmain.goをそのままdest.goにコピーしているだけである。
必要に応じて、整形処理を追加する。
*/
func (e *ExecutorGolang) Arrange() error {
	sourceCode, err := os.ReadFile(e.sourceCodePath.MainPath)
	if err != nil {
		return fmt.Errorf("failed to read source code: %v", err)
	}

	destFile, err := os.Create(e.destGoPath)
	if err != nil {
		return fmt.Errorf("failed to create dest file: %v", err)
	}

	_, err = destFile.WriteString(string(sourceCode))
	if err != nil {
		return fmt.Errorf("failed to write source code: %v", err)
	}

	return nil
}

func (e *ExecutorGolang) Compile() error {
	cmd := exec.Command("bash", "-c", "cd "+e.outputDirPath+" && go build -o dest")
	_, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to compile: %v", err)
	}

	return nil
}

func (e *ExecutorGolang) Execute(reader io.Reader, writer io.Writer, errorWriter io.Writer) error {
	cmd := exec.Command(e.outputDirPath + "/dest")
	cmd.Stdin = reader
	cmd.Stdout = writer
	cmd.Stderr = errorWriter

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute: %v", err)
	}

	return nil
}

func (e *ExecutorGolang) ArrangedFile() (*os.File, error) {
	return os.Open(e.destGoPath)
}
