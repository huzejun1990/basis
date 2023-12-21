// @Author huzejun 2023/12/13 5:24:00
package _case

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

func PoolCase() {
	target := "192.168.31.108"
	pool, err := GetPool(target)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 10; i++ {
		conn := &Conn{
			ID:     int64(i),
			Target: target,
			Status: ON,
		}
		pool.Put(conn)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				conn := pool.Get()
				fmt.Println(conn.ID)
				pool.Put(conn)
			}

		}()
	}
	wg.Wait()
}

const (
	ON  = 1
	OFF = 0
)

type Conn struct {
	ID     int64
	Target string
	Status int
}

func NewConn(target string) *Conn {
	return &Conn{
		ID:     rand.Int63(),
		Target: target,
		Status: ON,
	}
}

func (c *Conn) GetStatus() int {
	return c.Status
}

type ConnPool struct {
	sync.Pool
}

func GetPool(target string) (*ConnPool, error) {
	return &ConnPool{
		Pool: sync.Pool{
			New: func() any {
				return NewConn(target)
			},
		},
	}, nil

}

func (c *ConnPool) Get() *Conn {
	conn := c.Pool.Get().(*Conn)
	if conn.GetStatus() == OFF {
		conn = c.Pool.New().(*Conn)
	}
	return conn
}

func (c *ConnPool) Put(conn *Conn) {
	if conn.GetStatus() == OFF {
		return
	}
	c.Pool.Put(conn)
}

// get
// put
