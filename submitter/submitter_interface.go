package submitter

import "os"

// ソースコード提出に関する機能を提供する
type ISubmitter interface {
	/* ソースコードを提出する
	file: 提出するソースコードのファイル
	*/
	Submit(file *os.File) error
}
