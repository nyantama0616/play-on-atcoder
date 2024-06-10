package submitter

import "os"

// ソースコード提出に関する機能を提供する
type ISubmitter interface {
	/* ソースコードを提出する
	language: 提出する言語
	file: 提出するソースコードのファイル
	*/
	Submit(language string, file *os.File) error
}
