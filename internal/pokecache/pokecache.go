package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	pokemap map[string]cacheEntry
	mutex   sync.RWMutex
	Done    chan bool
}

func GetCurrentTime() string {
	currentTime := time.Now()
	return fmt.Sprintf("%02d:%02d:%02d", currentTime.Hour(), currentTime.Minute(), currentTime.Second())
}

func NewCache(interval time.Duration) *Cache {
	var newCache = Cache{pokemap: make(map[string]cacheEntry)}
	newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, ok := c.pokemap[key]
	if ok {
		return
	}
	c.pokemap[key] = cacheEntry{createdAt: time.Now(), val: value}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, ok := c.pokemap[key]
	if !ok {
		fmt.Println("Not cached")
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval * time.Second)
	go func() {
		for {
			select {
			case <-c.Done:
				ticker.Stop()
				return
			case <-ticker.C:
				fmt.Println("\n\n" + GetCurrentTime())
				fmt.Println("Updating cache")
				currentTime := time.Now()
				c.mutex.Lock()
				for key := range c.pokemap {
					difference := currentTime.Sub(c.pokemap[key].createdAt)
					if difference > interval*time.Second {
						fmt.Printf("\nDeleting: %s (%v Elapsed)", key, difference)
						delete(c.pokemap, key)
					}
				}
				c.mutex.Unlock()
				fmt.Print("\nPokedex > ")
			}
		}
	}()
}

func (c *Cache) Display() {
	fmt.Println(GetCurrentTime())
	fmt.Printf("Entries in cache: %d\n", len(c.pokemap))
	for key := range c.pokemap {
		entryTime := c.pokemap[key].createdAt
		fmt.Printf("%s: %02d:%02d:%02d\n", key, entryTime.Hour(), entryTime.Minute(), entryTime.Second())
	}
}
