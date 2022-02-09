package filter

import (
	"embed"
	"encoding/json"
	"io"
	"net/http"
)

//go:embed words.json
var f embed.FS

func Embedded() (Wordlist, error) {
	c, err := f.ReadFile("words.json")
	if err != nil {
		return nil, err
	}

	ret, err := read(c)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// this will download the wordlist from github/mdanka/szozat/main/src/constants/hungarian-word-letter-list.json.
func Download() (Wordlist, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/mdanka/szozat/main/src/constants/hungarian-word-letter-list.json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret, err := read(b)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func read(c []byte) (Wordlist, error) {
	var words Wordlist
	err := json.Unmarshal(c, &words)
	if err != nil {
		return nil, err
	}
	return words, nil
}
