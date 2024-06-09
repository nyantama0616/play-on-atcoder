package submitter

import (
	"fmt"
	"testing"

	"github.com/nyantama0616/play-on-atcoder/executor/cpp"
	"github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/session"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestSubmit(t *testing.T) {
	problem, _ := problem.NewProblem("abc354_a")
	session := session.NewSession()

	executor := cpp.NewExecutorCpp(problem, cpp.SourceCodePath{
		MainPath:       fmt.Sprintf("%s/executor/cpp/assets/main.cpp", setting.RootDir),
		IncludeDirPath: fmt.Sprintf("%s/executor/cpp/assets/include", setting.RootDir),
	})
	submitter := NewSubmitter(problem, session)

	executor.Arrange()
	sourceFile, _ := executor.ArrangedFile()
	defer sourceFile.Close()

	err := submitter.Submit(sourceFile)
	if err != nil {
		t.Errorf("Submit() should return nil, but got %v", err)
	}
}
