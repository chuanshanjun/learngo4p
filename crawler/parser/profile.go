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

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
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
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(contents, url, name)
				},
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
