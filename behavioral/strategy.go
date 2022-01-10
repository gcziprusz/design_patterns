package main

import "fmt"

type evictionStrategy interface {
	evict(c *cache)
}
type lifo struct{}
type fifo struct{}
type lru struct{}

func (l *lifo) evict(c *cache) {
	fmt.Println("Evict using the lifo strategy")
}
func (f *fifo) evict(c *cache) {
	fmt.Println("Evict using the fifo strategy")
}
func (r *lru) evict(c *cache) {
	fmt.Println("Evict using the lru strategy")
}

type cache struct {
	strategy    evictionStrategy
	storage     map[string]string
	capacity    int
	maxCapacity int
}

func initChache(strategy evictionStrategy) *cache {
	return &cache{
		strategy:    strategy,
		storage:     make(map[string]string, 3),
		capacity:    0,
		maxCapacity: 3,
	}
}
func (c *cache) setStrategy(strategy evictionStrategy) {
	c.strategy = strategy
}
func (c *cache) add(k, v string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	fmt.Printf("Adding %s %s \n", k, v)
	c.storage[k] = v
}
func (c *cache) evict() {
	c.strategy.evict(c)
	c.capacity--
}

func main() {
	lifoStrategy := &lifo{}
	fifoStrategy := &fifo{}
	lruStrategy := &lru{}

	cacheImpl := initChache(lifoStrategy)
	cacheImpl.add("a", "1")
	cacheImpl.add("b", "2")
	cacheImpl.add("c", "3")

	// eviction triggered
	cacheImpl.add("x", "0")

	cacheImpl.setStrategy(fifoStrategy)
	// eviction triggered using different strategy
	cacheImpl.add("xx", "00")

	cacheImpl.setStrategy(lruStrategy)
	// eviction triggered using different strategy
	cacheImpl.add("xxx", "000")
}
