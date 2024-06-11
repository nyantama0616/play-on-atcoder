package fetcher

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

type Fetcher struct {
	problem             problem.IProblem
	sampleNum           int
	sampleContainerPath string
}

// FetcherがIFetcherを実装していることを確認
var _ IFetcher = (*Fetcher)(nil)

/*
新しいFetcherを作成する

problem: 問題の情報を扱う構造体
*/
func NewFetcher(problem problem.IProblem) *Fetcher {
	sampleContainerPath := fmt.Sprintf("%s/fetcher/samples", problem.ProblemDirPath())
	if err := os.MkdirAll(sampleContainerPath, 0755); err != nil {
		panic(err)
	}

	return &Fetcher{
		problem:             problem,
		sampleNum:           0,
		sampleContainerPath: sampleContainerPath,
	}
}

func (f *Fetcher) Problem() problem.IProblem {
	return f.problem
}

/*
問題ページからサンプルケースを取得し、ファイルに保存する

	既にサンプルケースを取得している場合は何もしない
*/
func (f *Fetcher) FetchSamples() error {
	// 既にサンプルケースを取得している場合は何もしない
	if f.hasAlreadyFetched() {
		count := 0
		filepath.Walk(f.sampleContainerPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				count++
			}
			return nil
		})

		f.sampleNum = count - 1
		return nil
	}

	// 問題ページのURLを取得
	url := f.problem.ProblemUrl()

	/* HTTPリクエストを送信してHTMLを取得
	リトライ処理を行っている
	*/
	var resp *http.Response
	var err error
	for i := 0; i < setting.APIMaxRetry; i++ {
		resp, err = http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			break
		}
		if resp != nil {
			defer resp.Body.Close()
		}
		time.Sleep(setting.APIRetryInterval)
	}
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	// 「入力例」を含むh3要素の数を取得
	f.sampleNum = doc.Find("h3").FilterFunction(func(i int, s *goquery.Selection) bool {
		return strings.Contains(s.Text(), "入力例")
	}).Length()

	// 入力例を取得
	for i := 1; i <= f.sampleNum; i++ {
		// サンプルケースのディレクトリを作成
		sampleDirPath := f.sampleDirPath(i)
		if err := os.Mkdir(sampleDirPath, 0755); err != nil {
			return err
		}

		// 入力例のテキストを取得
		inputText := doc.Find(fmt.Sprintf("h3:contains('入力例 %d')", i)).Next().Text()

		// 入力例のテキストをファイルに書き込む
		inputFilePath := f.sampleInputFilePath(i)
		if err := os.WriteFile(inputFilePath, []byte(inputText), 0644); err != nil {
			return err
		}
	}

	// 出力例を取得
	for i := 1; i <= f.sampleNum; i++ {
		// 出力例のテキストを取得
		outputText := doc.Find(fmt.Sprintf("h3:contains('出力例 %d')", i)).Next().Text()

		// 出力例のテキストをファイルに書き込む
		outputFilePath := f.sampleOutputFilePath(i)
		if err := os.WriteFile(outputFilePath, []byte(outputText), 0644); err != nil {
			return err
		}
	}

	return nil
}

// Fetchしたサンプルケースの数を返す
func (f *Fetcher) SampleNum() int {
	return f.sampleNum
}

// サンプルiの入力ファイルを返す
func (f *Fetcher) SampleInputFile(i int) (*os.File, error) {
	file, err := os.Open(f.sampleInputFilePath(i))
	if err != nil {
		return nil, err
	}

	return file, nil
}

// サンプルiの出力ファイルを返す
func (f *Fetcher) SampleOutputFile(i int) (*os.File, error) {
	file, err := os.Open(f.sampleOutputFilePath(i))
	if err != nil {
		return nil, err
	}

	return file, nil
}

/*
サンプルケースのディレクトリのパスを返す

	例: contests/abc000/a/case1
*/
func (f *Fetcher) sampleDirPath(i int) string {
	return fmt.Sprintf("%s/case%d", f.sampleContainerPath, i)
}

/*
サンプルケースの入力ファイルのパスを返す

	例: contests/abc000/a/case1/input.txt
*/
func (f *Fetcher) sampleInputFilePath(i int) string {
	return fmt.Sprintf("%s/input.txt", f.sampleDirPath(i))
}

/*
サンプルケースの出力ファイルのパスを返す

	例: contests/abc000/a/case1/output.txt
*/
func (f *Fetcher) sampleOutputFilePath(i int) string {
	return fmt.Sprintf("%s/output.txt", f.sampleDirPath(i))
}

/*
既にサンプルケースを取得しているかを返す

	既にサンプルケースを取得している場合はtrue、そうでない場合はfalseを返す
*/
func (f *Fetcher) hasAlreadyFetched() bool {
	_, err := os.Stat(f.sampleInputFilePath(1))
	return err == nil
}
