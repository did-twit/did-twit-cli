package storage

import (
	"log"

	"github.com/tidwall/buntdb"
)

const (
	didIndex = "dids"
)

type DB struct {
	db *buntdb.DB
}

func NewConnection() (*DB, error) {
	db, err := buntdb.Open("storage.db")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.CreateIndex(didIndex, "did:twit:*", buntdb.IndexString); err != nil {
		return nil, err
	}
	return &DB{db: db}, err
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Write(key, value string) error {
	return db.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
}

func (db *DB) Read(key string) (*string, error) {
	var res string
	err := db.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		res = val
		return nil
	})
	return &res, err
}

func (db *DB) ListDIDs() ([]string, error) {
	var dids []string
	err := db.db.View(func(tx *buntdb.Tx) error {
		_ = tx.Ascend(didIndex, func(key, val string) bool {
			dids = append(dids, key)
			return true
		})
		return nil
	})
	return dids, err
}
