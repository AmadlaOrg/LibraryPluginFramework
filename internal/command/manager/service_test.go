package manager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPluginService(t *testing.T) {
	t.Run("should return a new instance of Plugin service", func(t *testing.T) {
		service := NewPluginService("plugin", "/home/user/.service")
		assert.NotNil(t, service)
		assert.IsType(t, &SManager{}, service)
		assert.Equal(t, service.GetPluginsPath(), "/home/user/.service/plugin")
	})
}
