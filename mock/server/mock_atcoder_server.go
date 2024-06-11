package server

import (
	"fmt"
	"net/http"

	"github.com/nyantama0616/play-on-atcoder/problem"
	"github.com/nyantama0616/play-on-atcoder/setting"
)

// AtCoderのサイトを模したサーバー
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
	handler.HandleFunc("/contests/abc354/submit", handlePostSubmit)
	handler.HandleFunc("/contests/abc354/submissions/me", handlePostSubmissionsMe)

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

			<form action="/contests/abc354/submit" method="POST">
				<select name="data.LanguageId">
					<option value="5001">C++ 20 (gcc 12.2)</option>
				</select>
				<input name="csrf_token" value="csrf_token">
			</form>
	</body>
</html>
		`))
}

func handlePostSubmit(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contests/abc354/submissions/me", http.StatusSeeOther)
}

func handlePostSubmissionsMe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
