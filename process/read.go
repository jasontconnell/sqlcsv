package process

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jasontconnell/sqlcsv/data"
	_ "github.com/microsoft/go-mssqldb"
)

func Read(connstr string, query string) (data.Table, error) {
	tbl := data.Table{}
	conn, err := sql.Open("mssql", connstr)
	if err != nil {
		return tbl, fmt.Errorf("couldn't open connection. %w", err)
	}
	defer conn.Close()

	rows, err := conn.Query(query)
	if err != nil {
		return tbl, fmt.Errorf("error running query %s. %w", query, err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return tbl, fmt.Errorf("couldn't get columns from query. %w", err)
	}

	for _, h := range cols {
		tbl.Headers = append(tbl.Headers, data.Column{
			Value: h,
		})
	}

	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}
	for rows.Next() {
		scerr := rows.Scan(vals...)
		if scerr != nil {
			log.Println("error reading row. ", scerr)
		}

		srow := []data.Column{}
		for _, v := range vals {
			deref := *(v.(*interface{}))
			srow = append(srow, data.Column{Value: fmt.Sprintf("%v", deref)})
		}
		tbl.Rows = append(tbl.Rows, data.Row{Columns: srow})
	}

	return tbl, nil
}
