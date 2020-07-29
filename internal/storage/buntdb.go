package storage

import (
	"log"

	"github.com/tidwall/buntdb"
)

type DB struct {
	*buntdb.DB
}

func NewConnection() (*DB, error) {
	db, err := buntdb.Open("storage.db")
	if err != nil {
		log.Fatal(err)
	}
	return &DB{DB: db}, err
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func (db *DB) Write(key, value string) error {
	return db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
}

func (db *DB) Read(key string) (*string, error) {
	var res string
	err := db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		res = val
		return nil
	})
	return &res, err
}
