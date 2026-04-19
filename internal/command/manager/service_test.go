package manager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should return a new instance of Plugin service", func(t *testing.T) {
		service := New("plugin", "/home/user/.service")
		assert.NotNil(t, service)
		assert.IsType(t, &managerImpl{}, service)
		assert.Equal(t, service.GetPluginsPath(), "/home/user/.service/plugin")
	})
}
