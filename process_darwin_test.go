// +build darwin

package ps

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindProcessDarwin(t *testing.T) {
	testFindProcess(t, "go-ps.test")
}

func TestProcessesDarwin(t *testing.T) {
	testProcesses(t, "go")
}

func TestProcessesDarwinError(t *testing.T) {
	errFn := func() ([]Process, error) {
		return nil, fmt.Errorf("oops")
	}
	proc, err := findProcessWithFn(errFn, os.Getpid())
	assert.Nil(t, proc)
	assert.EqualError(t, err, "Error listing processes: oops")
}

func TestProcessExecRun(t *testing.T) {
	testExecRun(t)
}

// TODO: Figure out how to get processes that have been removed
// func TestProcessExecRemoved(t *testing.T) {
// 	procPath, proc := testExecRun(t)
// 	// Remove it while it is running
// 	os.Remove(procPath)
// 	matchPath := func(p Process) bool { return p.Pid() == proc.Pid() }
// 	procs, err := findProcessesWithFn(processes, matchPath, 1)
// 	require.NoError(t, err)
// 	t.Logf("Proc: %#v", procs[0])
// }
