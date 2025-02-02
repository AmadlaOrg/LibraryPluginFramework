package manager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewManagerService(t *testing.T) {
	t.Run("should return a new instance of Plugin service", func(t *testing.T) {
		service := NewManagerService("plugin", "/home/user/.service")
		assert.NotNil(t, service)
		assert.IsType(t, &SManager{}, service)
		assert.Equal(t, service.GetPluginsPath(), "/home/user/.service/plugin")
	})
}
