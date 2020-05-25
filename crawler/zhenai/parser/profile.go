package parser

import (
	"regexp"
	"strconv"

	"chuanshan.github.com/learngo4p/crawler/engine"

	"chuanshan.github.com/learngo4p/crawler/model"
)

var genderRe = regexp.MustCompile(`"genderString":"(.士)"`)

var nameRe = regexp.MustCompile(`"nickname":"([^"].)"`)

var ageRe = regexp.MustCompile(`"age":([^,]+)`)

var heightRe = regexp.MustCompile(`"heightString":"([0-9]+)cm`)

var weightRe = regexp.MustCompile(`"([0-9]+)kg"`)

var incomeRe = regexp.MustCompile(`"月收入:([0-9-]+)千"`)

var marriageRe = regexp.MustCompile(`"marriageString":"([^,]+)"`)

var educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)

var idUrlRe = regexp.MustCompile(`href="http://album.zhenai.com/u/([0-9].+)">`)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	profile.Gender = extractString(contents, *genderRe)
	age, err := strconv.Atoi(extractString(contents, *ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, *heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, *weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, *incomeRe)
	profile.Marriage = extractString(contents, *marriageRe)
	profile.Education = extractString(contents, *educationRe)

	//return engine.ParseResult{
	//	Items: []interface{}{profile},
	//}

	result := engine.ParseResult{
		Requests: []engine.Request{
			{
				Url: url,
				//ParserFunc: func(bytes []byte) engine.ParseResult {
				//	return ParseProfile(contents, url, name)
				//},
				// 此处把上面给包装了下
				//ParserFunc: ProfileParser(name),
				Parser: NewProfileParser(name),
			},
		},
		Items: []engine.Item{
			{
				Url:  url,
				Type: "zhenai",
				// 可以通过匹配Url
				Id:      extractString(contents, *idUrlRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re regexp.Regexp) string {
	matches := re.FindSubmatch(contents)
	if len(matches) >= 2 {
		return string(matches[1])
	} else {
		return ""
	}
}

//func ProfileParser(name string) engine.ParserFunc {
//	return func(c []byte, url string) engine.ParseResult {
//		return ParseProfile(c, url, name)
//	}
//}

// contents,url是系统都知道,name是我们打包在函数的闭包里面的
// 所以我们构造ProfileParser里面的name就是userName
// 之前函数式编程是自动打包的，现在我们得手动打包
type ProfileParser struct {
	userName string
}

// 系统传进来的contents, url, 再加上我们自己的username，就给parseProfile配齐参数咯
func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

// 然后再加上Serialize,就是可以序列化/反序列化了
func (p *ProfileParser) Serialize() (name string, args interface{}) {
	// 对方用ProfileParser来解析的时候，如果需要name的时候，再把第二个参数p.userName传送过去
	return "ProfileParser", p.userName
}

// 帮助函数，存一个name，然后生成一个函数
func NewProfileParser(name string) *ProfileParser {
	// 如果要改变name的值，需要使用指针接受者
	// 反之如果不改变name值的话，就用值接收者即可
	// 使用指针接受者，少一次拷贝，性能更好
	return &ProfileParser{
		userName: name,
	}
}
