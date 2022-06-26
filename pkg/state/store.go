package state

import "errors"

// Store is an interface to access any key-value database used for store shortened urls.
type Store interface {
	// SaveUnique save the given key and is intended to guarantee that the key is unique.
	SaveUnique(short, long string) error
	// Retrieve returns the long version of the short URL.
	Retrieve(short string) (string, error)
}

var (
	ErrShortAlreadyBeenUsed = errors.New("short url is already in use")
	ErrShortNotFound        = errors.New("short url not found")
)
