package executor

import (
	"io"
	"os"
)

/*
提出可能なソースコードを用意し、コンパイル、実行する機能を提供する
*/
type IExecutor interface {
	// 提出可能なソースコードを用意する
	Arrange() error

	// ソースコードをコンパイルし、実行可能にする
	Compile() error

	/* ソースコードを実行する
	reader: 入力元
	writer: 出力先
	errorWriter: エラー出力先
	*/
	Execute(reader io.Reader, writer io.Writer, errorWriter io.Writer) error

	// ソースコード
	ArrangedFile() (*os.File, error)
}
