# 環境
### OS
Linux or WSL or Mac OS
### Golang
version 1.22.3

# 使用方法

### 環境変数の設定
.envファイルをプロジェクトのルートに作成し、AtCoderのusernameとpasswordを記述してください(絶対に外部に漏らさないようにすること)
.envファイルの書き方は`.env.sample`を参考にしてください。

### テスト
以下により、すべてのパッケージが問題なく動作するか確認します。
```bash
$ bin/test
```

### 回答を作成
AtCoderの問題を選び、回答を`_main.go`に記述してください。

### 検証&提出
下記のコマンドを実行し、回答を検証&提出します。
```bash
//problem_idの例: abc100_a
$ go run validate_and_submit.go <problem_id>
```

# その他
各パッケージの説明は[こちら](https://github.com/nyantama0616/play-on-atcoder/blob/master/doc/package.md)
