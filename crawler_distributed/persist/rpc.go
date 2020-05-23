package persist

import (
	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// 此处我们就用指针，省掉一次值拷贝也好
// *ItemSaverService是一个指针接收者,save方法不是开在ItemSaverService类型上的
// 而是开在ItemSaverService指针类型上的
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	if err == nil {
		*result = "ok"
	}
	return err
}
