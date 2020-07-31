package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	db, err := NewConnection()
	assert.NoError(t, err)

	// Write and then read
	err = db.Write("test", "val")
	assert.NoError(t, err)

	val, err := db.Read("test")
	assert.NoError(t, err)

	assert.Equal(t, "val", *val)

	// Read missing
	val, err = db.Read("missing")
	assert.Error(t, err)
	assert.Empty(t, val)

	// Close and try to write after closed
	assert.NoError(t, db.Close())
	assert.Error(t, db.Write("key", "val"))
}