package duplicatedetector

import (
	"bytes"

	"github.com/bradfitz/gomemcache/memcache"
)

// checker wraps some utility methods to work with memcached as duplicate detector
type checker struct {
	cache      *memcache.Client
	prefix     string
	expiration int32
	value      []byte
}

// NewChecker returns a new instance of a Duplicate Detector
// It can be configured with a special prefix (useful for namespacing different apps)
// and a TTL (in seconds, either a relative time from now, up to 1 month, or an absolute
// Unix epoch time. Zero means the items have no expiration time)
func NewChecker(mc *memcache.Client, prefix string, ttl int32) *checker {
	return &checker{cache: mc, prefix: prefix, expiration: ttl, value: []byte("x")}
}

// getKeyFor prepends the prefix to the item key
func (c *checker) getKeyFor(id string) string {
	return c.prefix + "_" + id
}

// getItemFor returns an Item object ready to be stored in memcache
func (c *checker) getItemFor(id string) *memcache.Item {
	return &memcache.Item{Key: c.getKeyFor(id), Value: c.value, Expiration: c.expiration}
}

// Set will unconditionally add the current item to the cache, even if it's already there
func (c *checker) Set(id string) error {
	return c.cache.Set(c.getItemFor(id))
}

// Has will check if the item has been previously seen already
func (c *checker) Has(id string) bool {
	v, err := c.cache.Get(c.getKeyFor(id))
	if err != nil {
		return false
	}
	if bytes.Equal(v.Value, c.value) {
		return true
	}
	return false
}

// Delete will remove the item from the cache, allowing a new Item with the same key in
func (c *checker) Delete(id string) error {
	err := c.cache.Delete(c.getKeyFor(id))
	if err == nil || err == memcache.ErrCacheMiss {
		return nil
	}
	return err
}

// IsDuplicate checks if the ID has been seen before (true) or if it's the first time (false).
// This counts as a touch: the first time an ID is checked, it is added to the cache;
// the second time the same ID is checked, it is considered as a duplicate
func (c *checker) IsDuplicate(id string) bool {
	if nil != c.cache.Add(c.getItemFor(id)) {
		return true
	}
	return false
}