// @Author huzejun 2023/12/12 20:53:00
package _case

import (
	"fmt"
	"sync"
)

func MutexCase() {
	//singleRoutine()
	//multipleRoutine()
	//multipleSafeRoutine()
	multipleSafeRoutineByRWMutex()
}

// 单协程操作
func singleRoutine() {
	mp := make(map[string]int, 0)
	list := []string{"A", "B", "C", "D"}
	for i := 0; i < 20; i++ {
		for _, item := range list {
			_, ok := mp[item]
			if !ok {
				mp[item] = 0
			}
			mp[item] += 1
		}
	}
	fmt.Println(mp)
}

// 多协程操作，非协程安全
func multipleRoutine() {
	mp := make(map[string]int, 0)
	list := []string{"A", "B", "C", "D"}

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, item := range list {
				_, ok := mp[item]
				if !ok {
					mp[item] = 0
				}
				mp[item] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp)
}

// 互斥锁协程安全
func multipleSafeRoutine() {
	type safeMap struct {
		data map[string]int
		sync.Mutex
	}
	mp := safeMap{
		data:  make(map[string]int, 0),
		Mutex: sync.Mutex{},
	}
	list := []string{"A", "B", "C", "D"}

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mp.Lock()
			defer mp.Unlock()
			for _, item := range list {
				_, ok := mp.data[item]
				if !ok {
					mp.data[item] = 0
				}
				mp.data[item] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp.data)
}

type cache struct {
	data map[string]string
	sync.RWMutex
}

func newCache() *cache {
	return &cache{
		data:    make(map[string]string, 0),
		RWMutex: sync.RWMutex{},
	}
}

func (c *cache) Get(key string) string {
	c.RLock()
	defer c.RUnlock()
	value, ok := c.data[key]
	if ok {
		return value
	}
	return ""
}

func (c *cache) Set(key, value string) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

func multipleSafeRoutineByRWMutex() {
	c := newCache()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Set("name", "nick")
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(c.Get("name"))
		}()
	}
	wg.Wait()
}
