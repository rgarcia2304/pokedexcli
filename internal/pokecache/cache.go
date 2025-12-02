package pokecache

import(
	"time"
	"sync"
)


type Cache struct{
	centry map[string]cacheEntry
	mu sync.Mutex
	lifetime time.Duration
}


type cacheEntry struct{
	createdAt time.Time
	value []byte
}

func NewCache(interval time.Duration) *Cache{
	nmap := make(map[string]cacheEntry)
	cch := Cache{lifetime: interval, 
		centry: nmap}
	go cch.reapLoop()
	return &cch

}

func (c *Cache) Add(key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.centry[key] = cacheEntry{value: val, createdAt: time.Now()}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.centry[key]

	if !ok{
		return nil, false
	}

	return val.value, ok 
}

func (c *Cache) reapLoop(){

	ticker := time.NewTicker(c.lifetime)
	
	for range ticker.C{

		c.mu.Lock() //lock the cache for read 
		for key, val := range c.centry{
			comp := time.Now().Add(-c.lifetime)
			if val.createdAt.Before(comp){
				delete(c.centry, key)	
			}
		}
		c.mu.Unlock()
	}
}
