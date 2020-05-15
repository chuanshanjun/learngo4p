package frontend

import (
	"html/template"
	"os"
	"testing"

	"chuanshan.github.com/learngo4p/crawler/engine"

	"chuanshan.github.com/learngo4p/crawler/frontend/model"
	// 重命名包名
	common "chuanshan.github.com/learngo4p/crawler/model"
)

func TestTemplate(t *testing.T) {
	// must我们认为模版语法没错误
	template := template.Must(
		template.ParseFiles("template.html"))

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	// 可以补充学习
	page.Hints = 123
	item := engine.Item{
		Url:  "http://hanxiao",
		Type: "zhenai",
		Id:   "abcDDE",
		Payload: common.Profile{
			Name:      "涵笑",
			Gender:    "女士",
			Age:       45,
			Height:    159,
			Weight:    52,
			Income:    "5-8",
			Marriage:  "离异",
			Education: "高中及以下",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	// os.Stdout 往屏幕输出
	//err := template.Execute(os.Stdout, page)
	// 不往屏幕输出往template.test.html输出
	err = template.Execute(out, page)
	if err != nil {
		panic(err)
	}
}
