# OneTextAPI-Go

A simple OneText Api powered by Golang.

## 使用方法

#### 在 Vercel 调用

example:

```go
package api

import (
    onetext "github.com/XiaoMengXinX/OneTextAPI-Go"
    "net/http"
)

var o = onetext.New()

func init() {
    _, err := o.GetUrl("https://raw.githubusercontent.com/lz233/OneText-Library/master/OneText-Library.json")
    if err != nil {
        panic(err)
    }
}

func Handler(w http.ResponseWriter, r *http.Request) {
    o.Handler(w, r)
}
```

#### 在本地运行

example:

```go
package main

import (
    onetext "github.com/XiaoMengXinX/OneTextAPI-Go"
)

func main() {
    o, err := onetext.New().ReadFile("./OneText-Library.json")
    if err != nil {
        panic(err)
    }
    o.StartServer(8000)
}
```

## 性能测试

```shell
$ wrk -t8 -c1000 -d10s --latency http://127.0.0.1:8000
Running 10s test @ http://127.0.0.1:8000
  8 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    71.47ms   63.77ms 440.57ms   72.18%
    Req/Sec     1.69k   587.43     4.13k    71.41%
  Latency Distribution
     50%   42.76ms
     75%  117.38ms
     90%  165.96ms
     99%  240.63ms
  129317 requests in 10.10s, 46.69MB read
  Socket errors: connect 0, read 0, write 584, timeout 0
Requests/sec:  12802.13
Transfer/sec:      4.62MB
```

