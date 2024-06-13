package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nyantama0616/play-on-atcoder/executor/golang"
	"github.com/nyantama0616/play-on-atcoder/fetcher"
	"github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/session"
	"github.com/nyantama0616/play-on-atcoder/setting"
	"github.com/nyantama0616/play-on-atcoder/submitter"
	"github.com/nyantama0616/play-on-atcoder/validator"
)

func main() {
	// .envから環境変数を読み込む
	envFilePath := fmt.Sprintf("%s/.env", setting.RootDir)
	if err := godotenv.Load(envFilePath); err != nil {
		panic("No .env file found")
	}

	problemId := os.Args[1]
	problem, err := problem.NewProblem(problemId)
	if err != nil {
		panic(err)
	}
	defer problem.RemoveProblemDir()

	fetcher := fetcher.NewFetcher(problem)
	err = fetcher.FetchSamples()
	if err != nil {
		panic(err)
	}

	executor := golang.NewExecutorGolang(
		problem,
		golang.SourceCodePath{
			MainPath: setting.RootDir + "/_main.go",
		},
	)
	err = executor.Arrange()
	if err != nil {
		panic(err)
	}

	err = executor.Compile()
	if err != nil {
		panic(err)
	}

	validator := validator.NewValidator(fetcher, executor)

	acCount := 0
	for i := 0; i < fetcher.SampleNum(); i++ {
		ok, err := validator.Validate(i + 1)
		if err != nil {
			panic(err)
		}
		if ok {
			acCount++
			fmt.Printf("Sample%d is \033[32mAC\033[0m\n", i+1)
		} else {
			fmt.Printf("Sample%d is \033[31mWA\033[0m\n", i+1)
		}
	}

	if acCount == fetcher.SampleNum() {
		fmt.Println("All samples are AC")

		//ログイン
		session := session.NewSession()
		if !session.IsLoggedIn() {
			session.LoginWithEnv()
		}

		//提出
		submitter := submitter.NewSubmitter(problem, session)
		sourceCode, _ := executor.ArrangedFile()
		if sourceCode == nil {
			panic("source code is nil")
		}
		err = submitter.Submit("Go (go 1.20.6)", sourceCode)
		if err != nil {
			panic(err)
		}
		fmt.Println("Submitted!")
	} else {
		fmt.Println("Some samples are WA")
	}
}
