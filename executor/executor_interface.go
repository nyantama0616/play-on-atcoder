package executor

import (
	"io"
	"os"
)

type IExecutor interface {
	Arrange() error
	Compile() error
	Execute(reader io.Reader, writer io.Writer, errorWriter io.Writer) error
	ArrangedFile() (*os.File, error)
}
