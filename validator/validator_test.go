package validator

import (
	"fmt"
	"testing"

	"github.com/nyantama0616/play-on-atcoder/executor/golang"
	. "github.com/nyantama0616/play-on-atcoder/fetcher"
	"github.com/nyantama0616/play-on-atcoder/mock"
	"github.com/nyantama0616/play-on-atcoder/mock/server"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestValidate(t *testing.T) {
	server := server.NewAtcoderServer(mock.NewMockProblem())
	listen := server.Setup()
	defer listen.Close()

	problem := mock.NewMockProblem()
	defer problem.RemoveProblemDir()

	fetcher := NewFetcher(problem)
	err := fetcher.FetchSamples()
	if err != nil {
		t.Errorf("FetchSamples() failed: %v", err)
	}

	executor := golang.NewExecutorGolang(
		problem,
		golang.SourceCodePath{
			MainPath: fmt.Sprintf("%s/executor/golang/_assets/main.go", setting.RootDir),
		},
	)

	executor.Arrange()

	err = executor.Compile()
	if err != nil {
		t.Errorf("%v", err)
	}

	validator := NewValidator(fetcher, executor)

	t.Run("すべてのサンプルがACになる", func(t *testing.T) {
		for i := 0; i < fetcher.SampleNum(); i++ {
			ok, err := validator.Validate(i + 1)
			if err != nil {
				t.Errorf("Validate() failed: %v", err)
			}
			if !ok {
				t.Errorf("Validate() failed: sample%d is not AC", i+1)
			}
		}
	})
}
