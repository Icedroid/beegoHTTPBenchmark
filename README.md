
## beego HTTP Benchmark testcases fot origin method, json encode, redis op, mongodb op, log etc...

I use [beego](http://www.beego.me/ "") bee tool command like follow:

```go
$ bee api beegoHTTPBenchmark
```
It will create an app hello in my gopath, and then I add three GET request method for compare raw beego /object,

Use redis client is [redigo]("https://github.com/garyburd/redigo", "")

use beego 1.2.0

benchmark on my ubuntu computer:

Linux ice-vm 3.8.0-33-generic #48~precise1-Ubuntu SMP Thu Oct 24 16:28:06 UTC 2013 x86_64 x86_64 x86_64 GNU/Linux

Intel(R) Pentium(R) CPU G2020 @ 2.90GHz

memory 8G

| :----:  | :----:  |

| router        |  function |

| /object     | origin bee api genarate http get method |

| /log     | origin bee api genarate http get method, but add one row log.Printf() |

| /set     | set data to redis hash list |

| /get     | get data from redis hash list |

I run this beego app, then use [wrk](https://github.com/wg/wrk "") for HTTP benchmarking, I get result:

### beego original method
```go
$ wrk -t12 -c400 -d30s http://127.0.0.1:8081/object 
Running 30s test @ http://127.0.0.1:8081/object
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.39s   375.09ms   1.65s    92.90%
    Req/Sec     1.35k     2.56k   14.22k    88.50%
  479750 requests in 30.01s, 161.05MB read
Requests/sec:  15987.51
Transfer/sec:      5.37MB
```

### set data to redis hash list use redigo
```go
$ wrk -t12 -c400 -d30s http://127.0.0.1:8081/set
Running 30s test @ http://127.0.0.1:8081/set
12 threads and 400 connections
Thread Stats   Avg      Stdev     Max   +/- Stdev
  Latency   223.28ms  397.74ms   1.41s    89.81%
  Req/Sec   399.67    270.22     5.04k    81.43%
135244 requests in 29.99s, 24.51MB read
Requests/sec:   4509.29
Transfer/sec:    836.69KB
```

### get data from redis hash list use redigo
```go
$ wrk -t12 -c400 -d30s http://127.0.0.1:8081/get
Running 30s test @ http://127.0.0.1:8081/get
12 threads and 400 connections
Thread Stats   Avg      Stdev     Max   +/- Stdev
  Latency    87.87ms   88.84ms 740.17ms   93.32%
  Req/Sec   417.59    129.96     1.55k    80.07%
149451 requests in 30.00s, 33.35MB read
Requests/sec:   4981.43
Transfer/sec:      1.11MB
```

### beego original method but add one row: log.Printf()
$ wrk -t12 -c400 -d30s http://127.0.0.1:8081/log   
Running 30s test @ http://127.0.0.1:8081/log
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    75.02ms  152.16ms   1.27s    85.61%
    Req/Sec   510.83    185.38     1.17k    68.07%
  183585 requests in 30.01s, 61.63MB read
Requests/sec:   6117.78
Transfer/sec:      2.05MB


#### I Don't know the result for Requests/sec  why have so  different, Redis Operation used redigo client is so ugly slow. Log also slow...