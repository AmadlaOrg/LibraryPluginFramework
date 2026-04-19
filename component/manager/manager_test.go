package manager

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	svc := New(cmd, "plugin", "/tmp/storage", nil)

	assert.NotNil(t, svc)
	assert.Equal(t, "/tmp/storage/plugin", svc.GetPluginsPath())
}
