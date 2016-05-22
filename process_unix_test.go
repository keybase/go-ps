// +build linux

package ps

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnixProcess(t *testing.T) {
	var _ Process = new(UnixProcess)
}

func TestFindProcessUnix(t *testing.T) {
	proc := testFindProcess(t, "go-ps.test")
	assert.True(t, proc.PPid() > 0)
}

func TestProcessesUnixError(t *testing.T) {
	errFn := func() ([]Process, error) {
		return nil, fmt.Errorf("oops")
	}
	proc, err := findProcessWithFn(errFn, os.Getpid())
	assert.Nil(t, proc)
	assert.EqualError(t, err, "Error listing processes: oops")
}
