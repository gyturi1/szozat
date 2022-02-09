package filter

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const WordLength = 5

type M string

const (
	Green  M = ":"
	Orange M = "+"
	Gray   M = "-"
)

var ValidMarkers = []string{string(Gray), string(Orange), string(Green)}

// represent a marked letter in a known position. Please note that not all marker needs the position info.
type Marker struct {
	Letter   Letter
	Position int
	M
}

// Markers to serach for / filter by.
type Markers []Marker

var (
	parser = fmt.Sprintf(`([\\%s\\%s\\%s]{0,1}\p{Latin}{0,3})`, string(Green), string(Orange), string(Gray))
	re     = regexp.MustCompile(parser)
)

// Parse guesses which must be WordLength letter long, and all letter prefixed with a marker.
func ParseAll(ss []string) (Markers, error) {
	var ret Markers
	for _, s := range ss {
		m, err := Parse(s)
		if err != nil {
			return nil, err
		}
		ret = append(ret, m...)
	}
	return ret, nil
}

// Parse one guess and returns the narkers describing each letter.
func Parse(s string) ([]Marker, error) {
	log.Info().Str("method", "Parse()").Msg("")
	var ret []Marker
	submatches := re.FindAllStringSubmatch(s, WordLength)
	log.Debug().Str("submatches", fmt.Sprintf("%v", submatches)).Msg("")

	if len(submatches) != WordLength {
		return nil, fmt.Errorf("invalid guess length got: %v", submatches)
	}

	c := 0
	for _, m := range submatches {
		m, l, err := parseSubMatch(m[0])
		if err != nil {
			return nil, errors.Wrap(err, "parse submatch")
		}
		ret = append(ret, Marker{Letter: l, Position: c, M: m})
		c++
	}
	return ret, nil
}

// parse a single letter in the guess into a marker.
func parseSubMatch(sm string) (M, Letter, error) {
	if len(sm) < 2 {
		return "", "", fmt.Errorf("both marker and letter are mandatory")
	}

	m := sm[:1]
	mark, err := parseMarker(m)
	if err != nil {
		return "", "", errors.Wrap(err, "parse marker")
	}

	l := sm[1:]
	if !validLetter(l) {
		return "", "", fmt.Errorf("invalid letter: %s", l)
	}

	return mark, l, nil
}

// parseMarker parses a marker string.
func parseMarker(s string) (M, error) {
	rs := []rune(s)
	if len(rs) == 0 {
		return "", fmt.Errorf("no marker")
	}
	if len(rs) > 1 {
		return "", fmt.Errorf("invalid marker: %v, valid markers: %v", rs, ValidMarkers)
	}
	r := rs[0]
	switch {
	case M(r) == Green:
		return Green, nil
	case M(r) == Gray:
		return Gray, nil
	case M(r) == Orange:
		return Orange, nil
	default:
		return "", fmt.Errorf("unknown marker: %v valid markers: %v", r, ValidMarkers)
	}
}
