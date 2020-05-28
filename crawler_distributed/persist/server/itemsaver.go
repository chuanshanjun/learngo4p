package main

import (
	"flag"
	"fmt"

	"chuanshan.github.com/learngo4p/crawler_distributed/config"
	"chuanshan.github.com/learngo4p/crawler_distributed/persist"
	"chuanshan.github.com/learngo4p/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	err := serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	// log.Fatal 强制退出，写起来更加简单
	//log.Fatal(serveRpc(":1234", "dating_profile"))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// 所以我们此时要取他的地址
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
