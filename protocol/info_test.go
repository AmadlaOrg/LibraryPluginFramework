package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInfo(t *testing.T) {
	t.Run("valid JSON", func(t *testing.T) {
		data := []byte(`{"name":"doorman-vault","version":"1.0.0","supports":["amadla.org/entity/secret@^v1.0.0"],"description":"HashiCorp Vault"}`)
		info, err := ParseInfo(data)
		assert.NoError(t, err)
		assert.Equal(t, "doorman-vault", info.Name)
		assert.Equal(t, "1.0.0", info.Version)
		assert.Equal(t, []string{"amadla.org/entity/secret@^v1.0.0"}, info.Supports)
		assert.Equal(t, "HashiCorp Vault", info.Description)
	})

	t.Run("empty supports", func(t *testing.T) {
		data := []byte(`{"name":"test","version":"0.1.0","supports":[],"description":""}`)
		info, err := ParseInfo(data)
		assert.NoError(t, err)
		assert.Equal(t, "test", info.Name)
		assert.Empty(t, info.Supports)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		data := []byte(`not json`)
		_, err := ParseInfo(data)
		assert.Error(t, err)
	})

	t.Run("empty input", func(t *testing.T) {
		_, err := ParseInfo([]byte{})
		assert.Error(t, err)
	})
}

func TestInfo_SupportsEntity(t *testing.T) {
	info := Info{
		Supports: []string{
			"amadla.org/entity/application@^v1.0.0",
			"amadla.org/entity/secret@^v1.0.0",
		},
	}

	t.Run("supported entity", func(t *testing.T) {
		assert.True(t, info.SupportsEntity("amadla.org/entity/application@^v1.0.0"))
	})

	t.Run("unsupported entity", func(t *testing.T) {
		assert.False(t, info.SupportsEntity("amadla.org/entity/infrastructure@^v1.0.0"))
	})

	t.Run("empty string", func(t *testing.T) {
		assert.False(t, info.SupportsEntity(""))
	})
}
