package filter

import (
	"embed"
	"encoding/json"
	"io"
	"net/http"
)

const etag = "e84c73fec7a54cb65a2868ce93c55bd2f3d0652fad2ae8e3d2b48ef526556208"

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
	url := "https://raw.githubusercontent.com/mdanka/szozat/main/src/constants/hungarian-word-letter-list.json"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("If-None-Match", etag)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 304 {
		return nil, nil
	}

	b, err := io.ReadAll(res.Body)
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
