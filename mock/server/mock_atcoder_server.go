package server

import (
	"fmt"
	"net/http"

	"github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

// AtCoderの問題ページを模したサーバ
type IAtcoderServer interface {
	Problem() problem.IProblem
	Setup() *http.Server
}

type AtcoderServer struct {
	problem problem.IProblem
}

// AtcoderServerがIAtcoderServerを実装していることを確認
var _ IAtcoderServer = (*AtcoderServer)(nil)

// 新しいAtcoderServerを生成する
func NewAtcoderServer(problem problem.IProblem) *AtcoderServer {
	return &AtcoderServer{
		problem: problem,
	}
}

func (s *AtcoderServer) Problem() problem.IProblem {
	return s.problem
}

func (s *AtcoderServer) Setup() *http.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/problem", handleGetProblem)

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", setting.MockServerPort),
		Handler: handler,
	}

	go server.ListenAndServe()

	return server
}

func handleGetProblem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
<html>
	<body>
			<h3>入力例 1</h3>
			<pre>54</pre>
			<h3>出力例 1</h3>
			<pre>6</pre>
			
			<h3>入力例 2</h3>
			<pre>7</pre>
			<h3>出力例 2</h3>
			<pre>4</pre>

			<h3>入力例 3</h3>
			<pre>262144</pre>
			<h3>出力例 3</h3>
			<pre>19</pre>
	</body>
</html>
		`))
}
