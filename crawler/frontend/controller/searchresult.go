package controller

import (
	"context"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"chuanshan.github.com/learngo4p/crawler/engine"

	"chuanshan.github.com/learngo4p/crawler/frontend/model"

	"chuanshan.github.com/learngo4p/crawler/frontend/view"
	"gopkg.in/olivere/elastic.v5"
)

// TODO
// fill in query string
// support search button
// support paging
// add start page
// rewrite query string
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=男 Payload.Age:(<30)&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		// 万一出错了
		from = 0
	}

	// 输出流
	//fmt.Fprintf(w, "q=%s, from=%d", q, from)
	var page model.SearchResult
	// 因为要用到client所以需要h
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).Do(context.Background())
	result.Query = q
	if err != nil {
		return result, err
	}

	result.Hints = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom =
		result.Start - len(result.Items)
	result.NextFrom =
		result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	// $1就是上面表达式中()内容
	return re.ReplaceAllString(q, "Payload.$1:")
}
