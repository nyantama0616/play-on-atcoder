package validator

import (
	"fmt"
	"testing"

	. "github.com/nyantama0616/play-on-atcoder/executor/cpp"
	. "github.com/nyantama0616/play-on-atcoder/fetcher"
	. "github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestValidate(t *testing.T) {
	problem, _ := NewProblem("abc354_a")
	defer problem.RemoveProblemDir()

	fetcher := NewFetcher(problem)
	fetcher.FetchSamples()

	executorCpp := NewExecutorCpp(
		problem,
		SourceCodePath{
			MainPath:       fmt.Sprintf("%s/executor/cpp/assets/main.cpp", setting.RootDir),
			IncludeDirPath: fmt.Sprintf("%s/executor/cpp/assets/include", setting.RootDir),
		},
	)

	err := executorCpp.Compile()
	if err != nil {
		t.Errorf("Compile() failed: %v", err)
	}

	validator := NewValidator(fetcher, executorCpp)

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
