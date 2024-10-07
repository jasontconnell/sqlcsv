package process

import (
	"io"
	"strings"

	"github.com/jasontconnell/sqlcsv/data"
)

func Write(out io.Writer, headers bool, tbl data.Table) error {
	var crlf []byte = []byte("\n")
	if headers {
		vals := []string{}
		for _, v := range tbl.Headers {
			vals = append(vals, v.Value)
		}
		headline := getCsvLine(vals)
		out.Write([]byte(headline))
		out.Write(crlf)
	}

	for _, r := range tbl.Rows {
		vals := []string{}
		for _, h := range r.Columns {
			vals = append(vals, h.Value)
		}
		line := getCsvLine(vals)
		out.Write([]byte(line))
		out.Write(crlf)
	}

	return nil
}

func getCsvLine(vals []string) string {
	line := ""
	for _, x := range vals {
		if strings.Contains(x, ",") {
			x = "\"" + x + "\""
		}
		line += x + ","
	}
	line = strings.TrimRight(line, ",")
	return line
}
