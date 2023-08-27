package mothership_worker_server

import (
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/base/go/storage"
)

type Worker struct {
	db      *base.DB
	storage storage.Storage
}

func New(db *base.DB, storage storage.Storage) *Worker {
	return &Worker{
		db:      db,
		storage: storage,
	}
}
