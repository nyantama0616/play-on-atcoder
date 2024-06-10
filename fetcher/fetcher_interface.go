package fetcher

import (
	"os"

	"github.com/nyantama0616/play-on-atcoder/problem"
)

// 問題のサンプルケースに関する機能を提供する
type IFetcher interface {
	// 問題の情報を扱う構造体を返す
	Problem() problem.IProblem

	// 問題のサンプルケースを問題ページから取得し、ファイルに保存する
	FetchSamples() error

	// 問題のサンプル数を返す
	SampleNum() int

	// サンプルiの入力ファイルを返す
	SampleInputFile(int) (*os.File, error)

	// サンプルiの出力ファイルを返す
	SampleOutputFile(int) (*os.File, error)
}
