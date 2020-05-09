package persist

import (
	"context"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		// sniff是客户端用来维护集群的状态的，集群不跑在本机，跑在docker
		// 但我们现在es跑在docker内网，内网我们看不见，所以没办法sniff
		elastic.SetSniff(false))
	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	// %+v 打印结构体中的字段
	//fmt.Printf("%+v", resp)
	return resp.Id, nil
}
