package pkg

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"sort"

	"github.com/olekukonko/tablewriter"
)

type Parser interface {
	Parse(io.Reader) (content, error)
}

type content struct {
	header []string
	rows   [][]string
}

func Format(p Parser, r io.Reader, w io.Writer) error {
	c, err := p.Parse(r)
	if err != nil {
		return err
	}

	formatTable(c, w)

	return nil
}

type CSVParser struct {}

func (CSVParser) Parse(reader io.Reader) (content, error) {
	r := csv.NewReader(reader)

	header, err := r.Read()
	if err != nil {
		return content{}, err
	}

	rows, err := r.ReadAll()
	if err != nil {
		return content{}, err
	}

	return content{
		header: header,
		rows:   rows,
	}, nil
}

type JSONParser struct {}

func (JSONParser) Parse(reader io.Reader) (content, error) {
	r := json.NewDecoder(reader)

	var rows []map[string]string
	if err := r.Decode(&rows); err != nil {
		return content{}, err
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

	return content{
		header: headers,
		rows:   outputRows,
	}, nil
}

func formatTable(c content, w io.Writer) {
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
