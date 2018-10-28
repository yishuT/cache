package cache

type NonCache struct {
	resource Resource
}

func NewNonCache(resource Resource, sz int) *NonCache {
	return &NonCache{
		resource: resource,
	}
}

func (c *NonCache) Get(key string) interface{} {
	return c.resource.Get(key)
}

func (c *NonCache) Put(key string, val interface{}) error {
	return c.resource.Put(key, val)
}

// make sure cache is kind of resource
var _ Resource = &LRUCache{}