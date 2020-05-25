package engine

// 名字不一定要写
type ParserFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// 我们将原先Request中ParserFunc变更为Parser接口
type Request struct {
	Url    string
	Parser Parser
}

// 放到distributed中
//type SerializedParser struct {
//	functionName string
//	args         interface{}
//}

// contents, url 是系统本身会给我们的，所以不用序列化
// {"ParseCityList", nil}, {"ProfileParser", userName}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct {
}

func (n NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// 实现了Parser接口
type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// 使用工厂函数创造FuncParser
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

// 改成分布式爬虫后注释掉
//func NilParser([]byte) ParseResult {
//	return ParseResult{}
//}
