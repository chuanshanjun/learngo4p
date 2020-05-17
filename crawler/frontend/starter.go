package main

import (
	"net/http"

	"chuanshan.github.com/learngo4p/crawler/frontend/controller"
)

func main() {
	// 不是search是/,我们就来展示文件
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))
	// 目前服务器只处理/search的请求,此时如果来一个/css/style.css 他当然不知道怎么handle
	//http.Handle("/search", controller.SearchResultHandler{})
	http.Handle("/search", controller.CreateSearchResultHandler("crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
