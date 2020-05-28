package main

import (
	"fmt"
	"log"

	"chuanshan.github.com/learngo4p/crawler_distributed/config"
	"chuanshan.github.com/learngo4p/crawler_distributed/rpcsupport"
	"chuanshan.github.com/learngo4p/crawler_distributed/worker"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{}))
}
