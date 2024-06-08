package executor

import "io"

type IExecutor interface {
	Compile() error
	Execute(reader io.Reader, writer io.Writer, errorWriter io.Writer) error
}
