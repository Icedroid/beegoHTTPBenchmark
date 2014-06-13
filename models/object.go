package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"log"
)

var (
	Objects map[string]*Object

	RedisPool      *redis.Pool
	server         = "127.0.0.1:6379"
	objectHashName = "object"
)

type Object struct {
	ObjectId   string
	Score      int64
	PlayerName string
}

func init() {
	Objects = make(map[string]*Object)
	Objects["hjkhsbnmn123"] = &Object{"hjkhsbnmn123", 100, "astaxie"}
	Objects["mjjkxsxsaa23"] = &Object{"mjjkxsxsaa23", 101, "someone"}

	RedisPool = &redis.Pool{
		MaxIdle:     5,
		MaxActive:   10000,
		IdleTimeout: time.Duration(5) * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				log.Printf("Dail redis server %s %v", server, err)
				return nil, err
			}
			if _, err := c.Do("PING"); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
}

func AddOne(object Object) (ObjectId string) {
	object.ObjectId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Objects[object.ObjectId] = &object
	return object.ObjectId
}

func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAll() map[string]*Object {
	return Objects
}

func Update(ObjectId string, Score int64) (err error) {
	if v, ok := Objects[ObjectId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectId Not Exist")
}

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}

//Add data to redis to redis hash list
func AddOneToRedis(object Object) (ObjectId string) {
	conn := GetRedis()
	defer conn.Close()
	object.ObjectId = "icedroid" + strconv.FormatInt(time.Now().UnixNano(), 10)
	_, err := conn.Do("HMSET", redis.Args{}.Add(objectHashName).AddFlat(object)...)
	if nil != err {
		log.Printf("set object to redis hash get error %s", err)
	}
	return object.ObjectId
}

//get data from redis hash list
func GetObject() (object *Object, err error) {
	conn := GetRedis()
	defer conn.Close()
	object = new(Object)
	reply, err := redis.Values(conn.Do("HGETALL", objectHashName))
	if nil != err {
		log.Printf("set object to redis hash get error %s", err)
		return
	}
	err = redis.ScanStruct(reply, object)
	return
}

func GetRedis() redis.Conn {
	//	c, err := redis.Dial("tcp", server)
	//	if err != nil {
	//		log.Printf("Dail master redis server %s %v", server, err)
	//		return nil
	//	}
	//	return c
	return RedisPool.Get()
}
