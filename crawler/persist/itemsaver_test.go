package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"chuanshan.github.com/learngo4p/crawler/engine"

	"gopkg.in/olivere/elastic.v5"

	"chuanshan.github.com/learngo4p/crawler/model"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://hanxiao",
		Type: "zhenai",
		Id:   "abcDDE",
		Payload: model.Profile{
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

	// TODO 目前ES依赖外部环境，但在测试中不希望出现此情况
	// Try to start up es here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// save expected item
	err = Save(client, expected, index)
	// 测试里面去panic它
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.
		Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	t.Logf("resp source %s: ", resp.Source)

	var actual engine.Item
	// type RawMessage []byte 就是一个 byte切片
	// err不拿也没关系，错了也直接挂掉
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	fmt.Printf("actual is %+v", actual)

	actualProfile, err := model.FromJsonObj(actual.Payload)
	if err != nil {
		panic(err)
	}
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("Expected %v, but got %v",
			expected, actual)
	}
}
