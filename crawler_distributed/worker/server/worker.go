package main

import (
	"flag"
	"fmt"
	"log"

	"chuanshan.github.com/learngo4p/crawler_distributed/rpcsupport"
	"chuanshan.github.com/learngo4p/crawler_distributed/worker"
)

// usage是 --help的提示
var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	// port在Parse里面会改变值，所以需要使用指针
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
