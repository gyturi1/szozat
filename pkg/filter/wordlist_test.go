package filter

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestIsCacheFile(t *testing.T) {
	tests := []struct {
		name string
		n    string
		want bool
	}{
		{"Emtpy", "", false},
		{"No prefix", "adfadfgadf.cache", false},
		{"Wrong suffix", "szozat_adfgadfgafd.che", false},
		{"Good I", "szozat_adfgadfgafd.cache", true},
		{"Good II", "szozat_13415135413.cache", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCacheFile(tt.n); got != tt.want {
				t.Errorf("IsCacheFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEtagFromFileName(t *testing.T) {
	tests := []struct {
		name string
		fn   string
		want string
	}{
		{"Empty", "", ""},
		{"Not cache file", "afgafd.adfgaf", ""},
		{"Valid I", "szozat_123456789.cache", "123456789"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EtagFromFileName(tt.fn); got != tt.want {
				t.Errorf("EtagFromFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_write_amd_readCache(t *testing.T) {
	etag := "123456789"
	b := []byte("testbytes:-)")

	defer os.Remove(cacheFilePath(etag))

	err := writeCache(etag, b)
	if err != nil {
		t.Error(err)
	}
	got, err := readCache(etag)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, b) {
		t.Errorf("readCache() = %v, want %v", string(got), string(b))
	}
}

func Test_newestCachedFileContent(t *testing.T) {
	etag1 := "111"
	etag2 := "222"
	c1 := []byte("rtrtrwwrtrwtrtr")
	c2 := []byte("agafgfgfagqrtqrtqretqr")

	defer os.Remove(cacheFilePath(etag1))
	defer os.Remove(cacheFilePath(etag2))

	err := writeCache(etag1, c1)
	if err != nil {
		t.Error(err)
	}

	// i know i know but just this time :-)
	time.Sleep(100 * time.Millisecond)
	err = writeCache(etag2, c2)
	if err != nil {
		t.Error(err)
	}

	got, _, err := newestCachedFileContent()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, c2) {
		t.Errorf("newestCachedFileContent() = %v, want %v", string(got), string(c2))
	}
}
