package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-spatial/geom/encoding/wkb"
	_ "github.com/jinzhu/gorm"
)

func PrintQueryResults(results *sql.Rows) error {
	rows := results
	//rows, err := config.PostDB.Table(results).Rows()
	//if err != nil {
	//	return err
	//}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	var values []interface{}
	for rows.Next() {
		record := make(map[string]interface{})
		values = make([]interface{}, len(columns))
		for i := range columns {
			values[i] = new(interface{})
			record[columns[i]] = nil
		}

		if err := rows.Scan(values...); err != nil {
			return err
		}

		for i := range columns {
			switch val := (*(values[i].(*interface{}))).(type) {
			case []byte:
				record[columns[i]] = string(val)
			default:
				record[columns[i]] = val
			}
		}

		fmt.Printf("%v\n", record)
	}

	return nil
}
