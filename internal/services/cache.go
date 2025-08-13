package services

type cache struct {
	data map[string]string
	hits int
}

func newCache() *cache {
	return &cache{data: make(map[string]string)}
}

func (c *cache) get(key string) string {
	if val, ok := c.data[key]; ok {
		c.hits++
		return val
	}
	return ""
}
