package pkg

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseCSV(t *testing.T) {
	tests := []struct {
		input    string
		expected Content
	}{
		{
			input: `id,name,price
1,apple,15
2,banana,10
`,
			expected: Content{
				header: []string{"id", "name", "price"},
				rows: [][]string{
					{"1", "apple", "15"},
					{"2", "banana", "10"},
				}},
		},
	}

	for i, test := range tests {
		out, err := CSVParser{}.Parse(strings.NewReader(test.input))
		require.NoError(t, err, i)
		assert.Equal(t, test.expected, out, i)
	}
}

func Test_parseJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected Content
	}{
		{
			input: `[
	{"id": "1", "name": "apple", "price": "15"},
	{"id": "2", "name": "banana", "price": "10"}
]`,
			expected: Content{
				header: []string{"id", "name", "price"},
				rows: [][]string{
					{"1", "apple", "15"},
					{"2", "banana", "10"},
				}},
		},
	}

	for i, test := range tests {
		out, err := JSONParser{}.Parse(strings.NewReader(test.input))
		require.NoError(t, err, i)
		assert.Equal(t, test.expected, out, i)
	}
}

func Test_collectHeader(t *testing.T) {
	tests := []struct {
		input    []map[string]string
		expected []string
	}{
		{
			input: []map[string]string{
				{"id": "1", "name": "apple"},
				{"id": "1", "price": "15"},
			},
			expected: []string{"id", "name", "price"},
		},
	}

	for i, test := range tests {
		out := collectHeader(test.input)

		// Unfortunately we can't guarantee the order of the headers due to the usage of maps.
		sort.Strings(out)
		sort.Strings(test.expected)
		assert.Equal(t, test.expected, out, i)
	}
}
