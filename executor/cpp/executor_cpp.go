package cpp

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/nyantama0616/play-on-atcoder/executor"
	"github.com/nyantama0616/play-on-atcoder/problem"
)

type ExecutorCpp struct {
	outputDirPath  string
	sourceCodePath SourceCodePath
	destCppPath    string
}

// IExecutorを実装しているか確認
var _ executor.IExecutor = (*ExecutorCpp)(nil)

type SourceCodePath struct {
	MainPath       string
	IncludeDirPath string
}

// TODO: sourceCodePathにデフォルト値を設定したい
func NewExecutorCpp(problem problem.IProblem, sourceCodePath SourceCodePath) *ExecutorCpp {
	outputDirPath := fmt.Sprintf("%s/executor", problem.ProblemDirPath())
	destCppPath := fmt.Sprintf("%s/dest.cpp", outputDirPath)

	//フォルダが存在しない場合は作成
	os.MkdirAll(outputDirPath, 0777)

	return &ExecutorCpp{
		outputDirPath:  outputDirPath,
		sourceCodePath: sourceCodePath,
		destCppPath:    destCppPath,
	}
}

func (e *ExecutorCpp) Compile() error {
	err := e.arrange()
	if err != nil {
		return fmt.Errorf("failed to arrange: %v", err)
	}

	cmd := exec.Command("g++-14", "-std=gnu++20", "-Wall", "-Wextra", "-O2", "-DONLINE_JUDGE", "-I", e.sourceCodePath.IncludeDirPath, "-o", e.outputDirPath+"/dest", e.destCppPath)
	_, err = cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to compile: %v", err)
	}

	return nil
}

func (e *ExecutorCpp) Execute(reader io.Reader, writer io.Writer, errorWriter io.Writer) error {
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

func (e *ExecutorCpp) arrange() error {
	sourceCode, err := os.ReadFile(e.sourceCodePath.MainPath) // ファイル全体を読み込む
	if err != nil {
		return fmt.Errorf("failed to read source code: %v", err)
	}

	definesCode, err := os.ReadFile(e.sourceCodePath.IncludeDirPath + "/mylib/defines.h")
	if err != nil {
		return fmt.Errorf("failed to read defines code: %v", err)
	}

	macrosCode, err := os.ReadFile(e.sourceCodePath.IncludeDirPath + "/mylib/macros.h")
	if err != nil {
		return fmt.Errorf("failed to read macros code: %v", err)
	}

	// sourceCodeを書き換える
	modifiedSourceCode := strings.Replace(string(sourceCode), "#include <mylib/macros.h>", string(macrosCode), -1)
	modifiedSourceCode = strings.Replace(modifiedSourceCode, "#include <mylib/defines.h>", string(definesCode), -1)
	modifiedSourceCode = strings.Replace(modifiedSourceCode, "#include <mylib/macros.h>", "", -1)
	modifiedSourceCode = strings.Replace(modifiedSourceCode, "#pragma once", "", -1)
	modifiedSourceCode = strings.Replace(modifiedSourceCode, "#define DEBUG_MODE 1", "#define DEBUG_MODE 0", -1)

	destFile, err := os.Create(e.destCppPath)
	if err != nil {
		return fmt.Errorf("failed to create dest file: %v", err)
	}

	_, err = destFile.WriteString(string(modifiedSourceCode))
	if err != nil {
		return fmt.Errorf("failed to write source code: %v", err)
	}

	return nil
}
