package runner

import "time"

const defaultTimeout = 30 * time.Second

// NewRunnerService creates a runner for the given plugin binary path.
//
// Example:
//
//	runner := New("/usr/local/bin/doorman-vault")
//	info, _ := runner.Info()
//	stdout, stderr, exitCode, _ := runner.Run(entityJSON, "get")
func New(pluginPath string) Runner {
	return &runnerImpl{
		pluginPath: pluginPath,
		timeout:    defaultTimeout,
	}
}

// NewRunnerServiceWithTimeout creates a runner with a custom timeout.
func NewWithTimeout(pluginPath string, timeout time.Duration) Runner {
	return &runnerImpl{
		pluginPath: pluginPath,
		timeout:    timeout,
	}
}
