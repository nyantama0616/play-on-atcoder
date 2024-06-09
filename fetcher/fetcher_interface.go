package fetcher

import (
	"os"

	"github.com/nyantama0616/play-on-atcoder/problem"
)

type IFetcher interface {
	Problem() problem.IProblem
	FetchSamples() error
	SampleNum() int
	SampleInputFile(int) (*os.File, error)
	SampleOutputFile(int) (*os.File, error)
}
