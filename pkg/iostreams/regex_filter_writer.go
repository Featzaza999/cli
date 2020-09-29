package iostreams

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
)

func NewRegexFilterWriter(out io.Writer, regexp *regexp.Regexp, repl string) io.Writer {
	return &RegexFilterWriter{out: out, regexp: *regexp, repl: repl}
}

type RegexFilterWriter struct {
	out    io.Writer
	regexp regexp.Regexp
	repl   string
}

func (s RegexFilterWriter) Write(data []byte) (int, error) {
	filtered := []byte{}
	repl := []byte(s.repl)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		b := scanner.Bytes()
		f := s.regexp.ReplaceAll(b, repl)
		if bytes.Equal(f, b) {
			f = append(f, []byte("\n")...)
		}
		filtered = append(filtered, f...)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	if len(filtered) == 0 {
		return len(data), nil
	}

	return s.out.Write(filtered)
}
