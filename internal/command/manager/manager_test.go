package manager

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPluginsPath(t *testing.T) {
	s := &managerImpl{pluginsPath: "/home/user/.service/plugin"}
	assert.Equal(t, "/home/user/.service/plugin", s.GetPluginsPath())
}

func TestList(t *testing.T) {
	s := &managerImpl{}
	result, err := s.List("", "")
	assert.NoError(t, err)
	assert.Equal(t, []string{}, result)
}

func TestAdd_NoCallback(t *testing.T) {
	s := &managerImpl{}
	err := s.Add([]string{"github.com/example/plugin@v1.0.0"})
	assert.NoError(t, err)
}

func TestAdd_WithCallback(t *testing.T) {
	var calledEvent EventType
	var calledAffected []Affected

	s := &managerImpl{
		runPluginManager: func(event EventType, affected []Affected, err error) error {
			calledEvent = event
			calledAffected = affected
			return nil
		},
	}

	err := s.Add([]string{"github.com/example/plugin@v1.0.0"})
	assert.NoError(t, err)
	assert.Equal(t, EventTypeAdd, calledEvent)
	assert.Len(t, calledAffected, 1)
	assert.Equal(t, "github.com/example/plugin@v1.0.0", calledAffected[0].Uri)
}

func TestAdd_CallbackError(t *testing.T) {
	s := &managerImpl{
		runPluginManager: func(event EventType, affected []Affected, err error) error {
			return errors.New("callback failed")
		},
	}

	err := s.Add([]string{"github.com/example/plugin"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "callback failed")
}

func TestRemove(t *testing.T) {
	s := &managerImpl{}
	err := s.Remove([]string{"github.com/example/plugin@v1.0.0"})
	assert.NoError(t, err)
}

func TestRemoveWithAllVersions(t *testing.T) {
	s := &managerImpl{}
	err := s.RemoveWithAllVersions([]string{"github.com/example/plugin"})
	assert.NoError(t, err)
}

func TestRemoveAll(t *testing.T) {
	s := &managerImpl{}
	err := s.RemoveAll()
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	s := &managerImpl{}
	err := s.Update([]string{"github.com/example/plugin"})
	assert.NoError(t, err)
}

func TestUpdateAll(t *testing.T) {
	s := &managerImpl{}
	err := s.UpdateAll()
	assert.NoError(t, err)
}

func TestUpdatable(t *testing.T) {
	s := &managerImpl{}
	err := s.Updatable([]string{"github.com/example/plugin"})
	assert.NoError(t, err)
}

func TestUpdatableAll(t *testing.T) {
	s := &managerImpl{}
	err := s.UpdatableAll()
	assert.NoError(t, err)
}

func TestSwitchVersion(t *testing.T) {
	s := &managerImpl{}
	err := s.SwitchVersion("github.com/example/plugin@v2.0.0", "github.com/example/plugin@v1.0.0")
	assert.NoError(t, err)
}
