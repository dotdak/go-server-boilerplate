package webserver

import (
	"log"

	"github.com/boltdb/bolt"
)

type Bolter struct {
	DB     *bolt.DB
	Path   string
	logger *log.Logger
}

func NewBolter(path string, logger *log.Logger) *Bolter {
	return &Bolter{
		Path:   path,
		logger: logger,
	}
}

func (b *Bolter) Open() error {
	db, err := bolt.Open(b.Path, 0600, nil)
	if err != nil {
		b.DB = nil
		return err
	}

	b.DB = db

	return nil
}
func (b *Bolter) Close() error {
	return b.DB.Close()
}
