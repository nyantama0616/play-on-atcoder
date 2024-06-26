package fetcher

import (
	"bufio"
	"fmt"
	"testing"

	"github.com/nyantama0616/play-on-atcoder/mock"
	"github.com/nyantama0616/play-on-atcoder/mock/server"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

func TestNewFetcher(t *testing.T) {
	t.Run("正しいproblemを渡すとFetcherが生成される", func(t *testing.T) {
		problem := mock.NewMockProblem()
		defer problem.RemoveProblemDir()

		fetcher := NewFetcher(problem)

		if fetcher == nil {
			t.Errorf("fetcher should not be nil")
		}
	})
}

/*
LookAtMe-4:

interfaceの活用例2: Mockを作成する

fetcherはAtCoderの問題ページからサンプルを取得するが、テストの度にAtCoderのサーバーにアクセスするのはあんまりよくない。
そこで、Atcoderサーバのモックを作成し、それを使ってテストを行う。

MockProblemは、本物のAtCoderの問題ページの情報ではなく、モックサーバーの情報を返す。
*/
func TestFetchSamples(t *testing.T) {
	server := server.NewAtcoderServer(mock.NewMockProblem())
	listen := server.Setup()
	defer listen.Close()

	problem := mock.NewMockProblem()
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
		fp, _ := fetcher.SampleInputFile(1)
		defer fp.Close()

		scanner := bufio.NewScanner(fp)

		scanner.Scan()
		inputText := scanner.Text()

		if inputText != "54" {
			t.Errorf("input should be 54, but got %s", inputText)
		}
	})

	t.Run("サンプル1の出力は6である", func(t *testing.T) {
		fp, _ := fetcher.SampleOutputFile(1)
		defer fp.Close()

		scanner := bufio.NewScanner(fp)

		scanner.Scan()
		inputText := scanner.Text()

		if inputText != "6" {
			t.Errorf("input should be 6, but got %s", inputText)
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

	t.Run("サンプル1の入力ファイルのパスが正しい", func(t *testing.T) {
		fp, _ := fetcher.SampleInputFile(1)
		defer fp.Close()

		expected := fmt.Sprintf("%s/contests/abc354/a/fetcher/samples/case1/input.txt", setting.RootDir)
		if fp.Name() != expected {
			t.Errorf("input file path should be %s, but got %s", expected, fp.Name())
		}
	})

	t.Run("サンプル1の出力ファイルのパスが正しい", func(t *testing.T) {
		fp, _ := fetcher.SampleOutputFile(1)
		defer fp.Close()

		expected := fmt.Sprintf("%s/contests/abc354/a/fetcher/samples/case1/output.txt", setting.RootDir)
		if fp.Name() != expected {
			t.Errorf("output file path should be %s, but got %s", expected, fp.Name())
		}
	})
}
