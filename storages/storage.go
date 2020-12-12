package storages

type StorageModule struct {
	Path string
	Type string
}

func NewStorage(path string, mode string) *StorageModule {
	storage := &StorageModule{
		Path: path,
		Type: mode,
	}

	return storage
}
