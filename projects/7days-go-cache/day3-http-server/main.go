package main

import (
	"fmt"
	"log"
	"net/http"

	"geecache"
)

// 使用 map 模拟了数据源 db
var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// curl http://localhost:8080/_geecache/scores/Tom
// 630

// curl http://localhost:8080/_geecache/scores/kkk
// kkk not exist

func main() {
	// 创建一个名为 scores 的 Group，若缓存为空，回调函数会从 db 中获取数据并返回
	geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "127.0.0.1:8080"
	peers := geecache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
