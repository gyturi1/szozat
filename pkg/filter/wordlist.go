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

	"github.com/rs/zerolog/log"
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

// returns the embedded wordlist
func Embedded() (Wordlist, string, error) {
	log.Info().Str("method", "Emdedded()").Msg("")
	log.Debug().Str("etag.const.embedded", etag).Msg("")
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

// returns the latest cached content from local filesystem.
func LatestCached() (Wordlist, string) {
	log.Info().Str("method", "LatestCached()").Msg("")
	b, etag, err := latestCachedFileContent()
	log.Debug().Str("etag", etag).Msg("")
	if err != nil {
		return nil, ""
	}
	c, err := unmarshal(b)
	if err != nil {
		return nil, ""
	}
	return c, etag
}

// this will download the wordlist from github/mdanka/szozat/main/src/constants/hungarian-word-letter-list.json and cache it into the local filesystem.
func Download(currentEtag string) (Wordlist, error) {
	log.Info().Str("method", "Download()").Msg("")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", wordListURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("If-None-Match", fmt.Sprintf(`"%s"`, currentEtag))
	log.Trace().Str("req", fmt.Sprintf("%v", req)).Msg("")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	et := res.Header.Get("etag")
	et = strings.ReplaceAll(et, "W/", "")
	et = strings.ReplaceAll(et, "\"", "")

	log.Debug().Str("res.statusCode", res.Status).Msg("")
	log.Debug().Str("res.etag", et).Msg("")
	log.Trace().Str("res", fmt.Sprintf("%v", res)).Msg("")

	if res.StatusCode == 304 {
		log.Debug().Msg("Readinng from cache")
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
	log.Info().Str("method", "unmarshall()").Msg("")
	var words Wordlist
	err := json.Unmarshal(c, &words)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func latestCachedFileContent() ([]byte, string, error) {
	log.Info().Str("method", "latestCachedFileContent()").Msg("")
	d := os.TempDir()
	log.Debug().Str("dir", d).Msg("")

	fis, err := ioutil.ReadDir(d)
	if err != nil {
		return nil, "", err
	}

	var t time.Time
	var fileName string

	for _, fi := range fis {
		n := fi.Name()
		if isCacheFile(n) && fi.ModTime().After(t) {
			fileName = n
			t = fi.ModTime()
			log.Debug().Str("filename", n).Msg("")
			log.Debug().Time("time", t).Msg("")
		}
	}

	b, err := os.ReadFile(path.Join(d, fileName))
	if err != nil {
		return nil, "", err
	}

	return b, etagFromFileName(fileName), nil
}

func isCacheFile(n string) bool {
	return strings.HasPrefix(n, cacheFilePrefix) && strings.HasSuffix(n, cacheFileSuffix)
}

func etagFromFileName(fn string) string {
	if !isCacheFile(fn) {
		return ""
	}
	woPrefix := strings.TrimPrefix(fn, cacheFilePrefix)
	log.Debug().Str("woPrefix", woPrefix).Msg("")
	return strings.TrimSuffix(woPrefix, cacheFileSuffix)
}

func writeCache(etag string, b []byte) error {
	log.Info().Str("method", "writeCache()").Msg("")
	return os.WriteFile(cacheFilePath(etag), b, 0o600)
}

func readCache(etag string) ([]byte, error) {
	log.Info().Str("method", "readCache()").Msg("")
	b, err := os.ReadFile(cacheFilePath(etag))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func cacheFilePath(etag string) string {
	log.Info().Str("method", "cacheFilePath()").Msg("")
	d := os.TempDir()
	log.Debug().Str("dir", d).Msg("")
	fp := path.Join(d, fmt.Sprintf(cacheFileNameFormat, etag))
	log.Debug().Str("filepath", fp).Msg("")
	return fp
}
