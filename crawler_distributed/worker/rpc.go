package worker

import "chuanshan.github.com/learngo4p/crawler/engine"

type CrawlService struct{}

// request的Parser 类型是interface 无法在网络上传播
func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
