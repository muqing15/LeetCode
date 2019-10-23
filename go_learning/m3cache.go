package go_learning

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Cache interface {
	Add(k string, v string)
	Update(k string, v string)
	Remove(k string)
	Get(k string) string
	Clean()
	Println()
	Size() int
}

type StringCache struct {
	lock  sync.RWMutex
	store map[string]string
}

func NewStringCache() Cache {
	return &StringCache{
		lock:  sync.RWMutex{},
		store: make(map[string]string, 0),
	}
}

func (s *StringCache) Add(k, v string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.store[k]; ok {
		fmt.Println(fmt.Sprintf("add failed,the key %s has already added\n", k))
		return
	}
	s.store[k] = v
}

func (s *StringCache) Update(k, v string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.store[k]; !ok {
		fmt.Println(fmt.Sprintf("update failed , the key %s is not existing\n", k))
		return
	}
	s.store[k] = v
}

func (s *StringCache) Remove(k string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.store[k]; !ok {
		fmt.Println(fmt.Sprintf("remove failed , the key %s is not existing\n", k))
		return
	}
	delete(s.store, k)
}

func (s *StringCache) Get(k string) string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if _, ok := s.store[k]; !ok {
		fmt.Println(fmt.Sprintf("failed get value , the key %s is not existing", k))
		return ""
	}
	return s.store[k]
}

func (s *StringCache) Clean() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.store = make(map[string]string)
}

func (s *StringCache) Println() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for k, v := range s.store {
		fmt.Println(fmt.Sprintf("%-6s : %s\n", k, v))
		return
	}
}

func (s *StringCache) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.store)
}

func M3Cache() {
	cache := NewStringCache()
	fmt.Println("add 3 items to cache ----")
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Add("key3", "value3")
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("get key2 data from cache ----")
	fmt.Println(cache.Get("key2"))
	fmt.Println("add another 3 items to cache ----")
	cache.Add("key4", "value4")
	cache.Add("key5", "value5")
	cache.Add("key6", "value6")
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("update key5 to value88888 ----")
	cache.Update("key5", "value88888")
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("update key9 to value99999 ----")
	cache.Update("key9", "value99999")
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("remove key5 item from cache ----")
	fmt.Println(cache.Get("key5"))
	cache.Remove("key5")
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("remove key7 item from cache ----")
	cache.Remove("key7")
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("clear the cache datas ----")
	cache.Clean()
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("add 100 datas to cache ----")
	for i := 0; i < 100; i++ {
		go cache.Add("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	}
	time.Sleep(5)
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()

}
