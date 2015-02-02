package main

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
)

// ResourceConn adapts a Redigo connection to a Vitess Resource.
type ResourceConn struct {
	redis.Conn
}

func (r ResourceConn) Close() {
	r.Conn.Close()
}

func main() {
	p := pools.NewResourcePool(func() (pools.Resource, error) {
			c, err := redis.Dial("tcp", ":6379")
			return ResourceConn{c}, err
		}, 100, 20000, time.Minute)
	defer p.Close()
	r, err := p.Get()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Put(r)
	c := r.(ResourceConn)
	n, err := c.Do("INFO")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("info=%s", n)
}
