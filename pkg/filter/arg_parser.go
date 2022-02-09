package filter

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

const WordLength = 5

type Marker string

const (
	Green  Marker = ":"
	Orange Marker = "+"
	Gray   Marker = "-"
)

var ValidMarkers = []string{string(Gray), string(Orange), string(Green)}

// represent a Mark marked letter
type Mark struct {
	Letter   Letter
	Position int
	Marker
}

// Pattern to search for
type Pattern []Mark

var (
	parser = fmt.Sprintf(`([\\%s\\%s\\%s]{0,1}\p{Latin}{0,3})`, string(Green), string(Orange), string(Gray))
	re     = regexp.MustCompile(parser)
)

// Parse guesses which must be WordLength letter long, and all letter prefixed with a marker.
func ParseAll(ss []string) (Pattern, error) {
	var ret Pattern
	for _, s := range ss {
		m, err := Parse(s)
		if err != nil {
			return nil, err
		}
		ret = append(ret, m...)
	}
	return ret, nil
}

func Parse(s string) ([]Mark, error) {
	var ret []Mark
	submatches := re.FindAllStringSubmatch(s, WordLength)

	if len(submatches) != WordLength {
		return nil, fmt.Errorf("invalid guess length got: %v", submatches)
	}

	c := 0
	for _, m := range submatches {
		m, l, err := parseSubMatch(m[0])
		if err != nil {
			return nil, errors.Wrap(err, "parse submatch")
		}
		ret = append(ret, Mark{Letter: l, Position: c, Marker: m})
		c++
	}
	return ret, nil
}

func parseSubMatch(sm string) (Marker, Letter, error) {
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
func parseMarker(s string) (Marker, error) {
	rs := []rune(s)
	if len(rs) == 0 {
		return "", fmt.Errorf("no marker")
	}
	if len(rs) > 1 {
		return "", fmt.Errorf("invalid marker: %v, valid markers: %v", rs, ValidMarkers)
	}
	r := rs[0]
	switch {
	case Marker(r) == Green:
		return Green, nil
	case Marker(r) == Gray:
		return Gray, nil
	case Marker(r) == Orange:
		return Orange, nil
	default:
		return "", fmt.Errorf("unknown marker: %v valid markers: %v", r, ValidMarkers)
	}
}
