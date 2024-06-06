package fetcher

type IFetcher interface {
	FetchSamples() error
	SampleNum() int
	SampleInput(int) (string, error)
	SampleOutput(int) (string, error)
}
