package execbc

import (
	"io"
	"os/exec"

	"github.com/AlexxIT/go2rtc/pkg/core"
)

type pipeCloser struct {
	io.Writer
	io.Closer
	cmd *exec.Cmd
}

func PipeCloser(cmd *exec.Cmd) (io.WriteCloser, error) {
	stdin, err := cmd.StdinPipe()

	if err != nil {
		return nil, err
	}
	return pipeCloser{stdin, stdin, cmd}, nil
}

func (p pipeCloser) Close() (err error) {
	return core.Any(p.Closer.Close(), p.cmd.Process.Kill(), p.cmd.Wait())
}
