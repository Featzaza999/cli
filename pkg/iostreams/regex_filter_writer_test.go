package iostreams

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Write(t *testing.T) {
	type input struct {
		in     string
		regexp *regexp.Regexp
		repl   string
	}
	type output struct {
		wantsErr bool
		out      string
		length   int
	}
	tests := []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "single line input",
			input: input{
				in:     "test",
				regexp: regexp.MustCompile("test"),
				repl:   "blah",
			},
			output: output{
				wantsErr: false,
				out:      "blah",
				length:   4,
			},
		},
		{
			name: "multiple line input",
			input: input{
				in:     "test",
				regexp: regexp.MustCompile("test"),
				repl:   "blah",
			},
			output: output{
				wantsErr: false,
				out:      "blah",
				length:   4,
			},
		},
		// {
		// 	name:   "no matches",
		// 	input:  input{},
		// 	output: output{},
		// },
		// {
		// 	name:   "no output",
		// 	input:  input{},
		// 	output: output{},
		// },
		// {
		// 	name:   "removes remote from git push output",
		// 	input:  input{},
		// 	output: output{},
		// },
		// {
		// 	name:   "handles scanner error",
		// 	input:  input{},
		// 	output: output{},
		// },
	}

	for _, tt := range tests {
		out := &bytes.Buffer{}
		writer := NewRegexFilterWriter(out, tt.input.regexp, tt.input.repl)
		t.Run(tt.name, func(t *testing.T) {
			length, err := writer.Write([]byte(tt.input.in))

			if tt.output.wantsErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.output.out, out.String())
			assert.Equal(t, tt.output.length, length)
		})
	}
}

// input2 := heredoc.Doc(`
// 	remote:
//	nomatch:
// 	remote: Create a pull request for 'regex-4' on GitHub by visiting:
// 	remote:      https://github.com/samcoe/cli-test/pull/new/regex-4
// 	remote:
// `)
