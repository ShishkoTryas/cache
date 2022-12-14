package cache

import (
	"errors"
	"fmt"
)

type cache struct {
	storage map[string]interface{}
}

func New() *cache {
	return &cache{
		storage: make(map[string]interface{}),
	}
}

func (c *cache) Set(key string, value interface{}) error {
	if len(key) == 0 {
		return errors.New("empty key")
	}
	c.storage[key] = value
	return nil
}

func (c cache) Get(key string) (interface{}, error) {
	if c.storage[key] == nil {
		return nil, fmt.Errorf("unknown key for get %s", key)
	}
	return c.storage[key], nil
}

func (c *cache) Delete(key string) error {
	if c.storage[key] == nil {
		return fmt.Errorf("unknown key for delete %s", key)
	}
	delete(c.storage, key)
	return nil
}
