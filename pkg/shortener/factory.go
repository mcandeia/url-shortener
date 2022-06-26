package shortener

import (
	"errors"
	"sync"
)

var (
	// engineInst keeps as a singleton the engine instance that will be used further.
	engineInst map[EngineID]Engine
	// factoryInit guard to keep initialization executing once.
	factoryInit sync.Once
)

var (
	// ErrFactoryNotInitialized is returned when tried to access engine without initialize factory first.
	ErrFactoryNotInitialized = errors.New("factory should be initialized first in order to use shortener engines")
	// ErrNotRecognizedEngine is returned when the engine is not recognized (i.e. does not exists).
	ErrNotRecognizedEngine = errors.New("the engine used is not recognized")
)

// InitFactory initializes the engine factories.
func InitFactory() {
	factoryInit.Do(func() {
		engineInst = map[EngineID]Engine{
			Base64: NewBase64(),
			Noop:   NewNoop(),
		}
	})
}

// Get receveis an EngineID and returns its implementation.
func Get(id EngineID) (Engine, error) {
	if engineInst == nil {
		return nil, ErrFactoryNotInitialized
	}

	eng, ok := engineInst[id]
	if !ok {
		return nil, ErrNotRecognizedEngine
	}

	return eng, nil
}
