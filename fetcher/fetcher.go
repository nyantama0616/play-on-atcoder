package fetcher

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nyantama0616/play-on-atcoder/problem"
)

type Fetcher struct {
	problem   problem.IProblem
	sampleNum int
}

// FetcherがIFetcherを実装していることを確認
var _ IFetcher = (*Fetcher)(nil)

func NewFetcher(problem problem.IProblem) *Fetcher {
	return &Fetcher{
		problem:   problem,
		sampleNum: 0,
	}
}

// 問題ページからサンプルケースを取得する
func (f *Fetcher) FetchSamples() error {
	// 既にサンプルケースを取得している場合は何もしない
	if f.hasAlreadyFetched() {
		// 末尾の改行を削除
		count := 0
		filepath.Walk(f.problem.ProblemDirPath(), func(path string, info os.FileInfo, err error) error {
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

	// HTTPリクエストを送信してHTMLを取得
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

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

func (f *Fetcher) SampleNum() int {
	return f.sampleNum
}

func (f *Fetcher) SampleInput(i int) (string, error) {
	input, err := os.ReadFile(f.sampleInputFilePath(i))
	if err != nil {
		return "", err
	}

	// 末尾の改行を削除
	inputStr := strings.TrimSuffix(string(input), "\n")

	return inputStr, nil
}

func (f *Fetcher) SampleOutput(i int) (string, error) {
	output, err := os.ReadFile(f.sampleOutputFilePath(i))
	if err != nil {
		return "", err
	}

	// 末尾の改行を削除
	outputStr := strings.TrimSuffix(string(output), "\n")

	return outputStr, nil
}

func (f *Fetcher) sampleDirPath(i int) string {
	return fmt.Sprintf("%s/case%d", f.problem.ProblemDirPath(), i)
}

func (f *Fetcher) sampleInputFilePath(i int) string {
	return fmt.Sprintf("%s/input.txt", f.sampleDirPath(i))
}

func (f *Fetcher) sampleOutputFilePath(i int) string {
	return fmt.Sprintf("%s/output.txt", f.sampleDirPath(i))
}

func (f *Fetcher) hasAlreadyFetched() bool {
	_, err := os.Stat(f.sampleInputFilePath(1))
	return err == nil
}
