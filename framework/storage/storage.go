package storage

// StorageType represents the type of storage to use.
type StorageType string

// Constants for different storage types.
const (
	StorageTypeFilesystem StorageType = "filesystem"
	StorageTypeCloud      StorageType = "cloud"
)

// persistence defines the interface for saving and loading data.
type persistence interface {
	save(data []byte) error
	load() ([]byte, error)
}

// Storage represents a storage object with configurable parameters.
type Storage struct {
	Type         StorageType
	ResourceName string
	Namespace    string
	Key          string
}

// NewStorage creates a new Storage instance with default values.
func NewStorage(resourceName, namespace, key string) *Storage {
	return &Storage{
		Type:         StorageTypeFilesystem,
		ResourceName: resourceName,
		Namespace:    namespace,
		Key:          key,
	}
}

// Save stores data using the configured storage mechanism.
func (s *Storage) Save(content []byte) error {
	storage := s.getStorage()

	return storage.save(content)
}

// Load retrieves data using the configured storage mechanism.
func (s *Storage) Load() ([]byte, error) {
	storage := s.getStorage()

	return storage.load()
}

// getStorage returns the appropriate storage implementation based on the provided type.
func (s *Storage) getStorage() persistence {
	switch s.Type {
	case StorageTypeCloud:
		return &cloudStorage{bucketName: s.ResourceName, directory: s.Namespace, name: s.Key}
	case StorageTypeFilesystem:
		fallthrough
	default:
		return &filesystemStorage{basePath: s.Namespace, name: s.Key}
	}
}
