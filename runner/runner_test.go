package runner

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func writeScript(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	err := os.WriteFile(path, []byte(content), 0755)
	assert.NoError(t, err)
	return path
}

func TestInfo(t *testing.T) {
	dir := t.TempDir()
	script := writeScript(t, dir, "test-plugin", `#!/bin/sh
case "$1" in
    info) echo '{"name":"test-plugin","version":"1.0.0","supports":["amadla.org/entity/test@^v1.0.0"],"description":"Test plugin"}' ;;
    *) echo "unknown" >&2; exit 2 ;;
esac
`)

	runner := New(script)
	info, err := runner.Info()

	assert.NoError(t, err)
	assert.Equal(t, "test-plugin", info.Name)
	assert.Equal(t, "1.0.0", info.Version)
	assert.Equal(t, []string{"amadla.org/entity/test@^v1.0.0"}, info.Supports)
	assert.Equal(t, "Test plugin", info.Description)
}

func TestInfo_NonZeroExit(t *testing.T) {
	dir := t.TempDir()
	script := writeScript(t, dir, "bad-plugin", `#!/bin/sh
echo "error message" >&2
exit 1
`)

	runner := New(script)
	_, err := runner.Info()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "exited with code 1")
}

func TestRun_StdinStdout(t *testing.T) {
	dir := t.TempDir()
	script := writeScript(t, dir, "echo-plugin", `#!/bin/sh
# Read stdin and echo it back
cat
`)

	runner := New(script)
	input := []byte(`{"_type":"test","_body":{"key":"value"}}`)
	stdout, stderr, exitCode, err := runner.Run(input, "process")

	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, input, stdout)
	assert.Empty(t, stderr)
}

func TestRun_Stderr(t *testing.T) {
	dir := t.TempDir()
	script := writeScript(t, dir, "stderr-plugin", `#!/bin/sh
echo "diagnostic message" >&2
echo '{"status":"pass"}'
`)

	runner := New(script)
	stdout, stderr, exitCode, err := runner.Run(nil, "validate")

	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, string(stdout), `"status":"pass"`)
	assert.Contains(t, string(stderr), "diagnostic message")
}

func TestRun_ExitCode1(t *testing.T) {
	dir := t.TempDir()
	script := writeScript(t, dir, "fail-plugin", `#!/bin/sh
echo '{"status":"fail"}'
exit 1
`)

	runner := New(script)
	stdout, _, exitCode, err := runner.Run(nil, "validate")

	assert.NoError(t, err) // Non-zero exit is NOT an execution error
	assert.Equal(t, 1, exitCode)
	assert.Contains(t, string(stdout), `"status":"fail"`)
}

func TestRun_ExitCode2(t *testing.T) {
	dir := t.TempDir()
	script := writeScript(t, dir, "usage-plugin", `#!/bin/sh
echo "Usage: plugin {info|validate}" >&2
exit 2
`)

	runner := New(script)
	_, stderr, exitCode, err := runner.Run(nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, exitCode)
	assert.Contains(t, string(stderr), "Usage:")
}

func TestRun_NonexistentPlugin(t *testing.T) {
	runner := New("/nonexistent/plugin")
	_, _, _, err := runner.Run(nil, "info")

	assert.Error(t, err)
}

func TestRunWithContext_Timeout(t *testing.T) {
	dir := t.TempDir()
	// Use a trap-resistant sleep so the process can be killed
	script := writeScript(t, dir, "slow-plugin", `#!/bin/sh
exec sleep 5
`)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	r := &runnerImpl{pluginPath: script, timeout: 200 * time.Millisecond}
	_, _, exitCode, err := r.RunWithContext(ctx, nil, "process")
	// Context timeout kills the process — either err is non-nil or exit code is non-zero
	if err == nil {
		assert.NotEqual(t, 0, exitCode)
	}
}

func TestNew(t *testing.T) {
	runner := New("/usr/local/bin/test-plugin")
	assert.NotNil(t, runner)
}

func TestNewWithTimeout(t *testing.T) {
	runner := NewWithTimeout("/usr/local/bin/test-plugin", 5*time.Second)
	assert.NotNil(t, runner)
}
