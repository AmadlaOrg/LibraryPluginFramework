package manager

import ()

type IManager interface {
	GetPluginsPath() string
	List(filterUri, filterVersion string) ([]string, error)
	Add(uri []string) error
	Remove(uri []string) error
	RemoveWithAllVersions(uri []string) error
	RemoveAll() error
	Update(uri []string) error
	UpdateAll() error
	Updatable(uri []string) error
	UpdatableAll() error
}

type SManager struct {
	placeholder      string
	storagePath      string
	pluginsPath      string
	runPluginManager RunPluginManager
}

// RunPluginManager
//
// Params:
// - 🎊 EventType - To inform the callback function on what event executed the call
// - 🧭 affectedAbsPath - The absolute path(s) of the plugin(s) that was/were affected
// - 🚨 error - If any errors occurred whilst the event happened
type RunPluginManager func(EventType, []Affected, error) error

// GetPluginsPath returns the complete absolute path to all the plugins
func (s *SManager) GetPluginsPath() string {
	return s.pluginsPath
}

// List
func (s *SManager) List(filterUri, filterVersion string) ([]string, error) {

	return []string{}, nil
}

// Add is for adding a plugin
func (s *SManager) Add(uri []string) error {
	// TODO: The version is taken from the URI... No version == latest

	s.runPluginManager(EventTypeAdd)

	return nil
}

// Remove only remove a specific plugin with the specific version
func (s *SManager) Remove(uri []string) error {

	return nil
}

// RemoveWithAllVersions remove all the versions of a specific plugin that exist in storage
func (s *SManager) RemoveWithAllVersions(uri []string) error {

	return nil
}

// RemoveAll goes into the storage path and deletes all the plugins
func (s *SManager) RemoveAll() error {

	return nil
}

// Update updates to latest version a specific plugin
func (s *SManager) Update(uri []string) error {

	return nil
}

// UpdateAll goes through all the plugins and adds the latest version and keeps the previous version
func (s *SManager) UpdateAll() error {

	return nil
}

// Updatable checks one plugin's version against its repo to verify if it is the latest version
func (s *SManager) Updatable(uri []string) error {
	return nil
}

// UpdatableAll checks each plugin versions against their respected repos to see which one has a new version
func (s *SManager) UpdatableAll() error {
	return nil
}

// SwitchVersion for when a change to specific version
func (s *SManager) SwitchVersion(uri, fromVersion string) error {
	// TODO: It pulls the specific version into the plugin storage
	// TODO: Then it removed the previous version
	err := s.Add(uri)
	if err != nil {
		return err
	}

	// TODO: Attach to URI the `fromVersion`

	return s.Remove(uri)
}
