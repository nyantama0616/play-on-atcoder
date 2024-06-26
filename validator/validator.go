package validator

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nyantama0616/play-on-atcoder/executor"
	"github.com/nyantama0616/play-on-atcoder/fetcher"
)

type Validator struct {
	fetcher   fetcher.IFetcher
	executor  executor.IExecutor
	outputDir string
}

// ValidatorがIValidatorを実装していることを確認
var _ IValidator = (*Validator)(nil)

/*
LookAtMe-4:

interfaceの活用例1: 挙動の変更を容易にする

interfaceという概念について詳しく学ぼう。
interfaceを使うことの最大のメリットは、「実装を隠蔽できる」ことである。

この説明での主要人物はvalidatorパッケージとexecutorパッケージである。
executorは、ソースコードをコンパイルし、実行するための機能を提供する。
ここで、ソースコードのコンパイル方法や実行方法はプログラミング言語によって異なる。
validatorはIExecutorに依存しているだけで、ソースコードのコンパイル方法や実行方法については知らない。
これにより、NewValidatorに渡すexecutorを変更するだけで、あらゆるプログラミング言語に対応できる。
*/
// 新しいValidatorを作成する
func NewValidator(fetcher fetcher.IFetcher, executor executor.IExecutor) *Validator {
	outputDir := fmt.Sprintf("%s/validator/answers", fetcher.Problem().ProblemDirPath())
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(err)
	}

	return &Validator{
		fetcher:   fetcher,
		executor:  executor,
		outputDir: outputDir,
	}
}

/*
問題のサンプルケースを実行し、出力を比較する
*/
func (v *Validator) Validate(num int) (bool, error) {
	input, err := v.fetcher.SampleInputFile(num)
	if err != nil {
		return false, err
	}
	defer input.Close()

	outputExpected, err := v.fetcher.SampleOutputFile(num)
	if err != nil {
		return false, err
	}
	defer outputExpected.Close()

	output, err := os.Create(fmt.Sprintf("%s/case%d.txt", v.outputDir, num))
	if err != nil {
		return false, err
	}
	defer output.Close()

	if err := v.executor.Execute(input, output, os.Stderr); err != nil {
		return false, err
	}
	output.Seek(0, 0) // ファイルポインタを先頭に戻す

	return v.compareOutputs(outputExpected, output), nil
}

// サンプルケースの出力と実行結果を比較する
func (v *Validator) compareOutputs(outputExpected *os.File, output *os.File) bool {
	scannerOutputExpected := bufio.NewScanner(outputExpected)
	scannerOutput := bufio.NewScanner(output)

	// TODO: スペースの違いを許容するようにする
	for scannerOutputExpected.Scan() {
		if !scannerOutput.Scan() {
			return false
		}

		if scannerOutputExpected.Text() != scannerOutput.Text() {
			return false
		}
	}

	return !scannerOutput.Scan()
}
