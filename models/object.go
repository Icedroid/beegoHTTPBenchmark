package models

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	r "menteslibres.net/gosexy/redis"
)

var (
	Objects map[string]*Object

	pool         *redis.Pool
	rPool, mPool *pools.ResourcePool
	session      *mgo.Session

	server      = "127.0.0.1:6379"
	maxIdle     = 10000
	idleTimeout = 24000

	url            = "127.0.0.1:27607"
	objectHashName = "object"

	client *r.Client
	host   = "127.0.0.1"
	port   = uint(6379)
)

type Object struct {
	ObjectId   string `bson:"_id,omitempty"`
	Score      int64
	PlayerName string
}

func init() {
	Objects = make(map[string]*Object)
	Objects["hjkhsbnmn123"] = &Object{"hjkhsbnmn123", 100, "astaxie"}
	Objects["mjjkxsxsaa23"] = &Object{"mjjkxsxsaa23", 101, "someone"}

	pool = newPool(server, maxIdle, idleTimeout)
	rPool = newResourcePool(server)
	mPool = newMongoResourcePool(url)
	session = newSession(url)
	//	client = GetRedis2(host, port)
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
func RAddOne(object Object) (ObjectId string) {
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
func RGetObject() (object *Object, err error) {
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

//get data from redis hash list
func RGetObject2() (object *Object, err error) {
	conn := GetRedis2()
	defer conn.Quit()
	//	object = new(Object)
	sl, err := conn.HMGet(objectHashName, "ObjectId", "Score", "PlayerName")
	if nil != err {
		log.Printf("get object from redis hash get error %s", err)
		return
	}
	//	log.Printf("sl=%v", sl)
	i, _ := strconv.ParseInt(sl[1], 10, 64)
	object = &Object{sl[0], i, sl[2]}
	//err = redis.ScanStruct(reply, object)
	return
}

//get data from redis hash list
func RGetObject3() (object *Object, err error) {
	conn := GetRedis3()
	defer CloseRedis(conn)
	object = new(Object)
	reply, err := redis.Values(conn.Do("HGETALL", objectHashName))
	if nil != err {
		log.Printf("set object to redis hash get error %s", err)
		return
	}
	err = redis.ScanStruct(reply, object)
	return
}

//Add data to mongo
func MAddOne(object Object) (ObjectId string) {
	conn := GetMongo()
	defer conn.Close()
	object.ObjectId = "icedroid" + strconv.FormatInt(time.Now().UnixNano(), 10)
	c := conn.DB("test").C("object")
	err := c.Insert(&object)
	if err != nil {
		log.Printf("set object to redis hash get error %s", err)
	}
	return object.ObjectId
}

//get one row from mongo
func MGetOne(ObjectId string) (object *Object, err error) {
	conn := GetMongo()
	defer conn.Refresh()
	object = new(Object)
	c := conn.DB("test").C("object")
	err = c.Find(bson.M{"_id": ObjectId}).One(&object)
	return
}

//get all row from mongo
func MGetObject() (objects []*Object, err error) {
	conn := GetMongo()
	defer conn.Close()
	object := new(Object)
	c := conn.DB("test").C("object")
	q := c.Find(nil)
	iter := q.Limit(10).Iter()
	//	iter.All(&objects)
	for iter.Next(object) {
		objects = append(objects, object)
	}
	err = iter.Close()
	return
}

//get one row from mongo
func MGetOne2(ObjectId string) (object *Object, err error) {
	conn := GetMongo2()
	defer CloseMongo(conn)
	object = new(Object)
	c := conn.DB("test").C("object")
	err = c.Find(bson.M{"_id": ObjectId}).One(&object)
	return
}

//get all row from mongo
func MGetObject2() (objects []*Object, err error) {
	conn := GetMongo2()
	defer CloseMongo(conn)
	object := new(Object)
	c := conn.DB("test").C("object")
	q := c.Find(nil)
	iter := q.Limit(10).Iter()
	//	iter.All(&objects)
	for iter.Next(object) {
		objects = append(objects, object)
	}
	err = iter.Close()
	return
}

// ResourceConn adapts a Redigo connection to a Vitess Resource.
type ResourceConn struct {
	redis.Conn
}

func (r ResourceConn) Close() {
	r.Conn.Close()
}

// ResourceConn adapts a Redigo connection to a Vitess Resource.
type ResourceMongoConn struct {
	*mgo.Session
}

func (r ResourceMongoConn) Close() {
	r.Close()
}

func newResourcePool(server string) *pools.ResourcePool {
	return pools.NewResourcePool(func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", server)
		return ResourceConn{c}, err
	}, 1, 10000, time.Minute)
}

func newMongoResourcePool(url string) *pools.ResourcePool {
	return pools.NewResourcePool(func() (pools.Resource, error) {
		return ResourceMongoConn{session.Clone()}, nil
	}, 1, 10000, time.Minute)
}

//create a redis pool
func newPool(server string, maxIdle, idleTimeout int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			//		  if _, err := c.Do("AUTH", password); err != nil {
			//			  c.Close()
			//			  return nil, err
			//		  }
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			//		  _, err := c.Do("PING")
			return nil
		},
	}
}

//create a mongo session
func newSession(url string) *mgo.Session {
	//	mgo.SetDebug(true)
	//	aLogger := log.New(os.Stdout, "mgo:", log.LstdFlags)
	//	mgo.SetLogger(aLogger)
	session, err := mgo.Dial(url)
	if err != nil {
		log.Printf("dail mongo err:%v", err)
		return nil
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})
	//	session.SetSyncTimeout(10 * time.Second)
	return session
}

func GetRedis() redis.Conn {
	//		c, err := redis.Dial("tcp", server)
	//		if err != nil {
	//			log.Printf("Dail master redis server %s %v", server, err)
	//			return nil
	//		}
	//		return c
	return pool.Get()
}

func GetRedis2() *r.Client {
	client := r.New()
	err := client.Connect(host, port)
	if err != nil {
		log.Fatalf("Connect failed: %s\n", err.Error())
	}
	return client
}

func GetRedis3() ResourceConn {
	r, err := rPool.Get()
	if err != nil {
		log.Fatal(err)
	}
	c := r.(ResourceConn)
	return c
}

func CloseRedis(r ResourceConn) {
	rPool.Put(r)
}

//获取mongo连接
func GetMongo() *mgo.Session {
	return session.Clone()
}

//获取mongo连接
func GetMongo2() ResourceMongoConn {
	r, err := mPool.Get()
	if err != nil {
		log.Fatal(err)
	}
	c := r.(ResourceMongoConn)
	return c
}

func CloseMongo(r ResourceMongoConn) {
	mPool.Put(r)
}
