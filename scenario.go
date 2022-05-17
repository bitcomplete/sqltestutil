package sqltestutil

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/jmoiron/sqlx"
)

// LoadScenario reads a YAML "scenario" file and uses it to populate the given
// db. Top-level keys in the YAML are treated as table names having repeated
// rows, where keys on each row are column names. For example:
//
//   users:
//	 - id: 1
//	   name: Alice
//	   email: alice@example.com
//	 - id: 2
//     name: Bob
//	   email: bob@example.com
//
//   posts:
//	 - user_id: 1
//	   title: Hello, world!
//	 - user_id: 2
//	   title: Goodbye, world!
//     is_draft: true
//
// The above would populate the users and posts tables. Fields that are missing
// from the YAML are left out of the INSERT statement, and so are populated with
// the default value for that column.
func LoadScenario(ctx context.Context, db sqlx.ExtContext, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	var result map[string][]map[string]interface{}
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	for table, rows := range result {
		for _, row := range rows {
			var columns []string
			var placeholders []string
			for column := range row {
				columns = append(columns, column)
				placeholders = append(placeholders, ":"+column)
			}
			sql := fmt.Sprintf(
				"INSERT INTO %q (%s) VALUES (%s)",
				table,
				strings.Join(columns, ", "),
				strings.Join(placeholders, ", "),
			)
			_, err = sqlx.NamedExecContext(ctx, db, sql, row)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
