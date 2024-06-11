package submitter

import (
	"fmt"
	"testing"

	"github.com/nyantama0616/play-on-atcoder/executor/cpp"
	"github.com/nyantama0616/play-on-atcoder/mock"
	"github.com/nyantama0616/play-on-atcoder/mock/server"
	"github.com/nyantama0616/play-on-atcoder/session"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestSubmit(t *testing.T) {
	server := server.NewAtcoderServer(mock.NewMockProblem())
	listen := server.Setup()
	defer listen.Close()

	problem := mock.NewMockProblem()
	// problem, _ := problem.NewProblem("abc354_a")
	defer problem.RemoveProblemDir()

	session := session.NewSession()

	executor := cpp.NewExecutorCpp(problem, cpp.SourceCodePath{
		MainPath:       fmt.Sprintf("%s/executor/cpp/assets/main.cpp", setting.RootDir),
		IncludeDirPath: fmt.Sprintf("%s/executor/cpp/assets/include", setting.RootDir),
	})
	submitter := NewSubmitter(problem, session)

	executor.Arrange()
	sourceFile, _ := executor.ArrangedFile()
	defer sourceFile.Close()
	language := "C++ 20 (gcc 12.2)"

	err := submitter.Submit(language, sourceFile)
	if err != nil {
		t.Errorf("Submit() should return nil, but got %v", err)
	}
}
