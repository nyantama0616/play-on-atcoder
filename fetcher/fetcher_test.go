package fetcher

import (
	"testing"

	"github.com/nyantama0616/play-on-atcoder/problem"
)

func TestNewFetcher(t *testing.T) {
	t.Run("正しいproblemを渡すとFetcherが生成される", func(t *testing.T) {
		problemId := "abc100_a"
		problem, _ := problem.NewProblem(problemId)
		defer problem.RemoveProblemDir()

		fetcher := NewFetcher(problem)

		if fetcher == nil {
			t.Errorf("fetcher should not be nil")
		}
	})
}

func TestFetchSamples(t *testing.T) {
	problemId := "abc354_a"
	problem, _ := problem.NewProblem(problemId)
	defer problem.RemoveProblemDir()

	fetcher := NewFetcher(problem)

	err := fetcher.FetchSamples()
	if err != nil {
		t.Errorf("err should be nil, but got %v", err)
	}

	t.Run("サンプルの数は3である", func(t *testing.T) {
		if fetcher.SampleNum() != 3 {
			t.Errorf("sampleNum should be 3, but got %d", fetcher.SampleNum())
		}
	})

	t.Run("サンプル1の入力は54である", func(t *testing.T) {
		input, _ := fetcher.SampleInput(1)

		if input != "54" {
			t.Errorf("input should be 54, but got %s", input)
		}
	})

	t.Run("サンプル1の出力は6である", func(t *testing.T) {
		output, _ := fetcher.SampleOutput(1)

		if output != "6" {
			t.Errorf("output should be 6, but got %s", output)
		}
	})

	t.Run("既にfetch済みの場合は、サンプル数のみを設定して何もしない", func(t *testing.T) {
		err := fetcher.FetchSamples()
		if err != nil {
			t.Errorf("err should be nil, but got %v", err)
		}

		if fetcher.SampleNum() != 3 {
			t.Errorf("sampleNum should be 3, but got %d", fetcher.SampleNum())
		}
	})
}
