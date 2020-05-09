package persist

import (
	"context"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"

	"chuanshan.github.com/learngo4p/crawler/model"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:      "涵笑",
		Gender:    "女士",
		Age:       45,
		Height:    159,
		Weight:    52,
		Income:    "5-8",
		Marriage:  "离异",
		Education: "高中及以下",
	}

	id, err := save(expected)
	// 测试里面去panic它
	if err != nil {
		panic(err)
	}

	// TODO 目前ES依赖外部环境，但在测试中不希望出现此情况
	// Try to start up es here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.
		Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	t.Logf("resp source %s: ", resp.Source)

	var actual model.Profile
	// type RawMessage []byte 就是一个 byte切片
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("Expected %v, but got %v",
			expected, actual)
	}
}
