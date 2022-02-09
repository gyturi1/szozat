package filter

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

const wordListURL = "https://raw.githubusercontent.com/mdanka/szozat/main/src/constants/hungarian-word-letter-list.json"

const (
	cacheFilePrefix     = "szozat_"
	cacheFileSuffix     = ".cache"
	cacheFileNameFormat = cacheFilePrefix + "%s" + cacheFileSuffix
)

// this is the etag of the embedded words.json NEEDS to be update if new version is embedded!
const etag = "e84c73fec7a54cb65a2868ce93c55bd2f3d0652fad2ae8e3d2b48ef526556208"

//go:embed words.json
var f embed.FS

func Embedded() (Wordlist, string, error) {
	c, err := f.ReadFile("words.json")
	if err != nil {
		return nil, "", err
	}

	ret, err := unmarshal(c)
	if err != nil {
		return nil, "", err
	}
	return ret, etag, nil
}

// this will download the wordlist from github/mdanka/szozat/main/src/constants/hungarian-word-letter-list.json.
func Download(currentEtag string) (Wordlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", wordListURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("If-None-Match", currentEtag)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	et := res.Header.Get("etag")
	et = strings.ReplaceAll(et, "W/", "")
	et = strings.ReplaceAll(et, "\"", "")

	println("StatusCode:" + res.Status)

	if res.StatusCode == 304 {
		cBytes, err := readCache(et)
		if err != nil {
			return nil, err
		}

		cached, err := unmarshal(cBytes)
		if err != nil {
			return nil, err
		}
		return cached, nil
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = writeCache(et, b)
	if err != nil {
		return nil, err
	}

	ret, err := unmarshal(b)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func unmarshal(c []byte) (Wordlist, error) {
	var words Wordlist
	err := json.Unmarshal(c, &words)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func LatestCached() (Wordlist, string) {
	b, etag, err := newestCachedFileContent()
	if err != nil {
		return nil, ""
	}
	c, err := unmarshal(b)
	if err != nil {
		return nil, ""
	}
	return c, etag
}

func newestCachedFileContent() ([]byte, string, error) {
	d := os.TempDir()
	fis, err := ioutil.ReadDir(d)
	if err != nil {
		return nil, "", err
	}

	var t time.Time
	var fileName string

	for _, fi := range fis {
		n := fi.Name()
		if IsCacheFile(n) && fi.ModTime().After(t) {
			fileName = n
			t = fi.ModTime()
		}
	}

	b, err := os.ReadFile(path.Join(d, fileName))
	if err != nil {
		return nil, "", err
	}

	return b, EtagFromFileName(fileName), nil
}

func IsCacheFile(n string) bool {
	return strings.HasPrefix(n, cacheFilePrefix) && strings.HasSuffix(n, cacheFileSuffix)
}

func EtagFromFileName(fn string) string {
	if !IsCacheFile(fn) {
		return ""
	}
	woPrefix := strings.TrimPrefix(fn, cacheFilePrefix)
	return strings.TrimSuffix(woPrefix, cacheFileSuffix)
}

func writeCache(etag string, b []byte) error {
	return os.WriteFile(cacheFilePath(etag), b, 0o600)
}

func readCache(etag string) ([]byte, error) {
	b, err := os.ReadFile(cacheFilePath(etag))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func cacheFilePath(etag string) string {
	d := os.TempDir()
	fp := path.Join(d, fmt.Sprintf(cacheFileNameFormat, etag))
	return fp
}
