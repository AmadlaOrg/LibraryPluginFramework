package manager

type Manager interface {
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

type managerImpl struct {
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
func (s *managerImpl) GetPluginsPath() string {
	return s.pluginsPath
}

// List
func (s *managerImpl) List(filterUri, filterVersion string) ([]string, error) {

	return []string{}, nil
}

// Add is for adding a plugin
func (s *managerImpl) Add(uri []string) error {
	if s.runPluginManager != nil {
		var affected []Affected
		for _, u := range uri {
			affected = append(affected, Affected{Uri: u})
		}
		if err := s.runPluginManager(EventTypeAdd, affected, nil); err != nil {
			return err
		}
	}

	return nil
}

// Remove only remove a specific plugin with the specific version
func (s *managerImpl) Remove(uri []string) error {

	return nil
}

// RemoveWithAllVersions remove all the versions of a specific plugin that exist in storage
func (s *managerImpl) RemoveWithAllVersions(uri []string) error {

	return nil
}

// RemoveAll goes into the storage path and deletes all the plugins
func (s *managerImpl) RemoveAll() error {

	return nil
}

// Update updates to latest version a specific plugin
func (s *managerImpl) Update(uri []string) error {

	return nil
}

// UpdateAll goes through all the plugins and adds the latest version and keeps the previous version
func (s *managerImpl) UpdateAll() error {

	return nil
}

// Updatable checks one plugin's version against its repo to verify if it is the latest version
func (s *managerImpl) Updatable(uri []string) error {
	return nil
}

// UpdatableAll checks each plugin versions against their respected repos to see which one has a new version
func (s *managerImpl) UpdatableAll() error {
	return nil
}

// SwitchVersion for when a change to specific version
func (s *managerImpl) SwitchVersion(uri, fromVersion string) error {
	if err := s.Add([]string{uri}); err != nil {
		return err
	}

	return s.Remove([]string{fromVersion})
}
