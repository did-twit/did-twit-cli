package pkg

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"reflect"
)

func GenerateEd25519Key() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(nil)
}

// Copy makes a 1-1 copy of src into dst.
func Copy(src interface{}, dst interface{}) error {
	if err := validateCopy(src, dst); err != nil {
		return err
	}
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dst)
}

func ToJSON(i interface{}) (string, error) {
	b, err := json.Marshal(i)
	return string(b), err
}

func validateCopy(src interface{}, dst interface{}) error {
	if src == nil {
		return errors.New("src is nil")
	}
	if dst == nil {
		return errors.New("dst is nil")
	}

	// Type check
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	if srcType != dstType {
		return errors.New("type of src and dst must match")
	}

	// Kind checks
	srcKind := srcType.Kind()
	if !(srcKind == reflect.Ptr || srcKind == reflect.Slice) {
		return errors.New("src is not of kind ptr or slice")
	}
	dstKind := dstType.Kind()
	if !(dstKind == reflect.Ptr || dstKind == reflect.Slice) {
		return errors.New("dst is not of kind ptr or slice")
	}
	return nil
}
