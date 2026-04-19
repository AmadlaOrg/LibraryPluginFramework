package runner

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/AmadlaOrg/LibraryPluginFramework/protocol"
)

// Runner executes plugin binaries as subprocesses.
type Runner interface {
	// Info calls "<plugin> info" and parses the JSON response.
	Info() (protocol.Info, error)

	// Run executes "<plugin> <args...>" with the given input on stdin.
	// Returns stdout bytes, stderr bytes, and exit code.
	Run(input []byte, args ...string) (stdout []byte, stderr []byte, exitCode int, err error)

	// RunWithContext executes with a context for timeout/cancellation.
	RunWithContext(ctx context.Context, input []byte, args ...string) (stdout []byte, stderr []byte, exitCode int, err error)
}

// runnerImpl implements Runner for a single plugin binary.
type runnerImpl struct {
	pluginPath string
	timeout    time.Duration
}

// For testing: allow overriding exec.Command
var execCommand = exec.CommandContext

// Info calls the plugin's info subcommand and parses the response.
func (s *runnerImpl) Info() (protocol.Info, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	stdout, stderr, exitCode, err := s.RunWithContext(ctx, nil, "info")
	if err != nil {
		return protocol.Info{}, fmt.Errorf("failed to execute %s info: %w", s.pluginPath, err)
	}
	if exitCode != 0 {
		return protocol.Info{}, fmt.Errorf("%s info exited with code %d: %s", s.pluginPath, exitCode, string(stderr))
	}

	return protocol.ParseInfo(stdout)
}

// Run executes the plugin with the given args and stdin input.
func (s *runnerImpl) Run(input []byte, args ...string) ([]byte, []byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	return s.RunWithContext(ctx, input, args...)
}

// RunWithContext executes the plugin with context for timeout/cancellation.
func (s *runnerImpl) RunWithContext(ctx context.Context, input []byte, args ...string) ([]byte, []byte, int, error) {
	cmd := execCommand(ctx, s.pluginPath, args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	if input != nil {
		cmd.Stdin = bytes.NewReader(input)
	}

	err := cmd.Run()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
			// Non-zero exit is not an execution error — it's the plugin's response
			err = nil
		}
	}

	return stdoutBuf.Bytes(), stderrBuf.Bytes(), exitCode, err
}
