package problem

/*
問題の情報を取得するための機能を提供する
*/
type IProblem interface {
	/*
		問題のIDを取得する
			例: "abc100_a"
	*/
	ProblemId() string

	/*
		コンテスト名を取得する
			例: "abc100"
	*/
	ContestName() string

	/*
		問題の難易度を取得する
			例: "a"
	*/
	Rank() string

	/*
		問題のディレクトリのパスを取得する
			例: "contests/abc100/a"
	*/
	ProblemDirPath() string

	/*
		問題のURLを取得する
			例: "https://atcoder.jp/contests/abc100/tasks/abc100_a"
	*/
	ProblemUrl() string

	/*
		提出ページのURLを取得する
			例: "https://atcoder.jp/contests/abc100/submissions/me"
	*/
	SubmissionUrl() string

	/*
		問題のディレクトリを作成する
			例: "contests/abc100/a"
	*/
	CreateProblemDir() error

	/*
		問題のディレクトリを削除する
			例: "contests/abc100/a"
	*/
	RemoveProblemDir() error
}
