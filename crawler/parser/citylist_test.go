package parser

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	expectUrl := [3]string{
		"http://m.zhenai.com/zhenghun/beijing",
		"http://m.zhenai.com/zhenghun/shanghai",
		"http://m.zhenai.com/zhenghun/tianjin",
	}

	//expectItem := [3]string{
	//	"City: 北京", "City: 上海", "City: 天津",
	//}

	result := ParseCityList(contents)
	const urlSize = 481
	if len(result.Requests) != urlSize {
		log.Printf("Expect get url size %d, but got %d", urlSize, len(result.Requests))
	}
	for i := 0; i < 2; i++ {
		if expectUrl[i] != result.Requests[i].Url {
			log.Printf("Expect get url %s, but got %s", expectUrl[i], result.Requests[i].Url)
		}
	}
	for i, url := range expectUrl {
		if result.Requests[i].Url != url {
			t.Errorf("Expect get utl %s, but got %s", url, result.Requests[i].Url)
		}
	}

	const itemSize = 481
	if len(result.Items) != itemSize {
		log.Printf("Expect get items size %d, but got %d", itemSize, len(result.Items))
	}
	//for i := 0; i < 2; i++ {
	//	if expectItem[i] != result.Items[i] {
	//		log.Printf("Expect get item %s, but got %s", expectItem[i], result.Items[i])
	//	}
	//}
	//for i, item := range expectItem {
	//	if result.Items[i] != expectItem[i] {
	//		t.Errorf("Expect get item %s, but got %s", item, result.Items[i])
	//	}
	//}
}
