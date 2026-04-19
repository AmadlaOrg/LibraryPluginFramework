package amadla

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	apps := map[string]string{"hery": "^1"}
	entities := map[string]string{"system": "v1"}

	svc := New(apps, entities)

	assert.NotNil(t, svc)
	assert.NotNil(t, svc.AmadlaCmd())
	assert.Equal(t, "amadla", svc.AmadlaCmd().Use)
}
