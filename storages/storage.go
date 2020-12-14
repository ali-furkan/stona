package storages

import "stona/tools/logger"

type StorageModule struct {
	Path string
	Type string
}

func NewStorage(path string, mode string) *StorageModule {
	storage := &StorageModule{
		Path: path,
		Type: mode,
	}

	logger.Debug("Storage", "{ path: "+path+" } '"+mode+"' Storage Created")

	return storage
}
