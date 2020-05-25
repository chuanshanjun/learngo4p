package client

import (
	"log"

	"chuanshan.github.com/learngo4p/crawler_distributed/config"

	"chuanshan.github.com/learngo4p/crawler_distributed/rpcsupport"

	"chuanshan.github.com/learngo4p/crawler/engine"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			// Call rpc
			result := ""
			// 我们此处在等待RPC的回复，但我们不用担心你，此处我们本就是在
			// goroutine中等待
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil || result != "ok" {
				log.Printf("Item saver: error saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}
