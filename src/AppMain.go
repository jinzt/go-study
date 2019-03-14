package main

import (
	"fmt"
	"go-study/src/document"
	"go-study/src/log"
	"go-study/src/redis"
)

func main() {
	fmt.Println("init main")
	document.DemoPropertiesV1()
	document.DemoPropertiesV2()
	log.DemoLog()
	redis.DemoRedis()
}
