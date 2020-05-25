package parser

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	const nameRe = `"nickname":"([^"].)"`

	re := regexp.MustCompile(nameRe)
	matches := re.FindAllSubmatch(contents, -1)

	const expectName = "涵笑"

	if expectName != string(matches[0][1]) {
		t.Errorf("Expect get name %s, but got %s", expectName, matches[0][1])
	}
}
