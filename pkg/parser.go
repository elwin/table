package pkg

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"sort"

	"github.com/olekukonko/tablewriter"
)

// Parser describes an interface to Parse an arbitrary document into
// our intermediate Content form.
type Parser interface {
	Parse(io.Reader) (Content, error)
}

// Content is the intermediate representation before it is converted
// to a table format.
type Content struct {
	header []string
	rows   [][]string
}

// Format converts the content of the reader to a table format using
// the supplied parser and writes it to the writer.
func Format(p Parser, r io.Reader, w io.Writer) error {
	c, err := p.Parse(r)
	if err != nil {
		return err
	}

	formatTable(c, w)

	return nil
}

// CSVParser is a parser implementation that parses CSV documents.
type CSVParser struct{}

// Parse converts the content of a reader to the Content representation.
func (CSVParser) Parse(reader io.Reader) (Content, error) {
	r := csv.NewReader(reader)

	header, err := r.Read()
	if err != nil {
		return Content{}, err
	}

	rows, err := r.ReadAll()
	if err != nil {
		return Content{}, err
	}

	return Content{
		header: header,
		rows:   rows,
	}, nil
}

// JSONParser is a parser implementation that parses JSON documents.
type JSONParser struct{}

// Parse converts the content of a reader to the Content representation.
func (JSONParser) Parse(reader io.Reader) (Content, error) {
	r := json.NewDecoder(reader)

	var rows []map[string]string
	if err := r.Decode(&rows); err != nil {
		return Content{}, err
	}

	headers := collectHeader(rows)
	sort.Strings(headers)
	var outputRows [][]string

	for _, row := range rows {
		outputRow := make([]string, len(headers))
		for i, header := range headers {
			outputRow[i] = row[header]
		}

		outputRows = append(outputRows, outputRow)
	}

	return Content{
		header: headers,
		rows:   outputRows,
	}, nil
}

func formatTable(c Content, w io.Writer) {
	table := tablewriter.NewWriter(w)
	table.SetHeader(c.header)
	table.AppendBulk(c.rows)
	table.Render()
}

func collectHeader(rows []map[string]string) []string {
	headerMap := map[string]struct{}{}
	for _, row := range rows {
		for k := range row {
			headerMap[k] = struct{}{}
		}
	}

	var out []string
	for header := range headerMap {
		out = append(out, header)
	}

	return out
}
