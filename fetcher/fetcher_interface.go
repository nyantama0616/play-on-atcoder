package fetcher

import "os"

type IFetcher interface {
	FetchSamples() error
	SampleNum() int
	SampleInputFile(int) (*os.File, error)
	SampleOutputFile(int) (*os.File, error)
}
