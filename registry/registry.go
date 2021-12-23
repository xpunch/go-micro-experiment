// Package registry is an interface for service discovery
package registry

// Default represents default registry instance
// can be replaced with different implementations
var Default Registry

// The registry provides an interface for service discovery
// and an abstraction over varying implementations
type Registry interface {
	Init(...Option) error
	Options() Options
	Register(*Service, ...RegisterOption) error
	Deregister(*Service, ...DeregisterOption) error
	GetService(string, ...GetOption) ([]*Service, error)
	ListServices(...ListOption) ([]*Service, error)
	Watch(...WatchOption) (Watcher, error)
	String() string
}

type Service struct {
	Name        string            `json:"name"`
	Nodes       []*Node           `json:"nodes,omitempty"`
	Handlers    []*Handler        `json:"endpoints,omitempty"`
	Subscribers []*Subscriber     `json:"subscribers,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type Node struct {
	Id       string            `json:"id"`
	Version  string            `json:"version,omitempty"`
	Address  string            `json:"address,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type Handler struct {
	Name     string            `json:"name"`
	Stream   bool              `json:"stream,omitempty"`
	Request  *Value            `json:"request,omitempty"`
	Response *Value            `json:"response,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type Subscriber struct {
	Topic    string            `json:"topic"`
	Payload  *Value            `json:"payload,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values,omitempty"`
}

// Register a service node. Additionally supply options such as TTL.
func Register(s *Service, opts ...RegisterOption) error {
	return Default.Register(s, opts...)
}

// Deregister a service node
func Deregister(s *Service) error {
	return Default.Deregister(s)
}

// Retrieve a service. A slice is returned since we separate Name/Version.
func GetService(name string) ([]*Service, error) {
	return Default.GetService(name)
}

// List the services. Only returns service names
func ListServices() ([]*Service, error) {
	return Default.ListServices()
}

// Watch returns a watcher which allows you to track updates to the registry.
func Watch(opts ...WatchOption) (Watcher, error) {
	return Default.Watch(opts...)
}

func String() string {
	return Default.String()
}
