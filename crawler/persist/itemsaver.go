package persist

import (
	"context"
	"errors"
	"log"

	"chuanshan.github.com/learngo4p/crawler/engine"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// sniff是客户端用来维护集群的状态的，集群不跑在本机，跑在docker
		// 但我们现在es跑在docker内网，内网我们看不见，所以没办法sniff
		elastic.SetSniff(false))
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

			err := Save(client, item, index)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error {
	// save一次创建一个client太烂了
	//client, err := elastic.NewClient(
	//	// sniff是客户端用来维护集群的状态的，集群不跑在本机，跑在docker
	//	// 但我们现在es跑在docker内网，内网我们看不见，所以没办法sniff
	//	elastic.SetSniff(false))
	//if err != nil {
	//	return err
	//}

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	// %+v 打印结构体中的字段
	//fmt.Printf("%+v", resp)
	return nil
}
