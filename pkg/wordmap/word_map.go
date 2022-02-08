package wordmap

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed words.json
var f embed.FS

var m map[string]struct{}
var empty struct{}

func Contains(w string) bool {
	_, ok := m[w]
	return ok
}

func init() {
	if m == nil {
		c, err := f.ReadFile("words.json")
		if err != nil {
			fmt.Println("can not init words map from words.json")
			return
		}

		m, err = Read(c)
		if err != nil {
			fmt.Println("can not init words map from words.json")
			return
		}
	}
}

func Read(c []byte) (map[string]struct{}, error) {
	var ret map[string]struct{} = make(map[string]struct{})
	var words [][]string

	err := json.Unmarshal(c, &words)
	if err != nil {
		return nil, err
	}

	for _, w := range words {
		ret[strings.Join(w, "")] = empty
	}

	return ret, nil
}
