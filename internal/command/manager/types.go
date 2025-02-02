package manager

// EventType
type EventType string

const (
	EventTypeList   EventType = "list"
	EventTypeAdd    EventType = "add"
	EventTypeUpdate EventType = "update"
)

// Affected contains the details of what was affected (no need to add "how" because it is pass in some other way)
type Affected struct {
	Uri            string
	Version        string
	AbsStoragePath string
}
